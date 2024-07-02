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

type CopyFileRequest struct {
	// SourceFilePath is the full path of the file you want to copy.
	//
	// For example - /path/to/file.jpg
	SourceFilePath string `json:"sourceFilePath"`
	// DestinationPath is the full path to the folder you want to copy the above file into.
	//
	// For example - /folder/to/copy/into/
	DestinationPath string `json:"destinationPath"`
}

//
// METHODS
//

// CopyFile will copy a file from one folder to another.
func (s *MediaService) CopyFile(ctx context.Context, r *CopyFileRequest) error {
	if r == nil {
		return errors.New("request is empty")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(r)
	if err != nil {
		return nil
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/copy", b, requestTypeAPI)
	if err != nil {
		return err
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")

	err = s.client.do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
