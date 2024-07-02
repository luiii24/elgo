// Copyright © 2023 whatsmeow
// Redeveloped by Amirul Dev

package appstate

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"hash"

	"github.com/amiruldev20/waSocket/appstate/lthash"
	waProto "github.com/amiruldev20/waSocket/binary/proto"
)

type Mutation struct {
	Operation waProto.SyncdMutation_SyncdOperation
	Action    *waProto.SyncActionValue
	Index     []string
	IndexMAC  []byte
	ValueMAC  []byte
}

type HashState struct {
	Version uint64
	Hash    [128]byte
}

func (hs *HashState) updateHash(mutations []*waProto.SyncdMutation, getPrevSetValueMAC func(indexMAC []byte, maxIndex int) ([]byte, error)) ([]error, error) {
	var added, removed [][]byte
	var warnings []error

	for i, mutation := range mutations {
		if mutation.GetOperation() == waProto.SyncdMutation_SET {
			value := mutation.GetRecord().GetValue().GetBlob()
			added = append(added, value[len(value)-32:])
		}
		indexMAC := mutation.GetRecord().GetIndex().GetBlob()
		removal, err := getPrevSetValueMAC(indexMAC, i)
		if err != nil {
			return warnings, fmt.Errorf("failed to get value MAC of previous SET operation: %w", err)
		} else if removal != nil {
			removed = append(removed, removal)
		} else if mutation.GetOperation() == waProto.SyncdMutation_REMOVE {
			// TODO figure out if there are certain cases that are safe to ignore and others that aren't
			// At least removing contact access from WhatsApp seems to create a REMOVE op for your own JID
			// that points to a non-existent index and is safe to ignore here. Other keys might not be safe to ignore.
			warnings = append(warnings, fmt.Errorf("%w for %X", ErrMissingPreviousSetValueOperation, indexMAC))
			//return ErrMissingPreviousSetValueOperation
		}
	}

	lthash.WAPatchIntegrity.SubtractThenAddInPlace(hs.Hash[:], removed, added)
	return warnings, nil
}

func uint64ToBytes(val uint64) []byte {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, val)
	return data
}

func concatAndHMAC(alg func() hash.Hash, key []byte, data ...[]byte) []byte {
	h := hmac.New(alg, key)
	for _, item := range data {
		h.Write(item)
	}
	return h.Sum(nil)
}

func (hs *HashState) generateSnapshotMAC(name WAPatchName, key []byte) []byte {
	return concatAndHMAC(sha256.New, key, hs.Hash[:], uint64ToBytes(hs.Version), []byte(name))
}

func generatePatchMAC(patch *waProto.SyncdPatch, name WAPatchName, key []byte, version uint64) []byte {
	dataToHash := make([][]byte, len(patch.GetMutations())+3)
	dataToHash[0] = patch.GetSnapshotMac()
	for i, mutation := range patch.Mutations {
		val := mutation.GetRecord().GetValue().GetBlob()
		dataToHash[i+1] = val[len(val)-32:]
	}
	dataToHash[len(dataToHash)-2] = uint64ToBytes(version)
	dataToHash[len(dataToHash)-1] = []byte(name)
	return concatAndHMAC(sha256.New, key, dataToHash...)
}

func generateContentMAC(operation waProto.SyncdMutation_SyncdOperation, data, keyID, key []byte) []byte {
	operationBytes := []byte{byte(operation) + 1}
	keyDataLength := uint64ToBytes(uint64(len(keyID) + 1))
	return concatAndHMAC(sha512.New, key, operationBytes, keyID, data, keyDataLength)[:32]
}
