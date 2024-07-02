package imagekit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
)

//
// REQUESTS
//

type DeleteFilesRequest struct {
	// FileIDs is the list of unique ID of the uploaded files.
	FileIDs []string `json:"fileIds"`
}

//
// RESPONSES
//

type DeleteFilesResponse struct {
	// SuccessfullyDeletedFileIDs is the array of fileIds which are successfully deleted.
	SuccessfullyDeletedFileIDs []string `json:"successfullyDeletedFileIds"`
}

//
// METHODS
//

// DeleteFiles deletes multiple files uploaded in media library using bulk file delete API.
//
// When you delete a file, all its transform are also deleted.
// However, if a file or specific transformation has been requested in the past, then the response is cached in CDN.
// You can purge the cache from the CDN using purge API.
func (s *MediaService) DeleteFiles(ctx context.Context, r *DeleteFilesRequest) (*DeleteFilesResponse, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}
	if len(r.FileIDs) == 0 {
		return nil, errors.New("file ids are empty")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(r)
	if err != nil {
		return nil, err
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/batch/deleteByFileIds", b, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Submit the request
	res := new(DeleteFilesResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
