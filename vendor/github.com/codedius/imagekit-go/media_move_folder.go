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

type MoveFolderRequest struct {
	// SourceFolderPath is the full path to the source folder you want to move.
	//
	// For example - /path/of/source/folder
	SourceFolderPath string `json:"sourceFolderPath"`
	// DestinationPath is the full path to the destination folder where you want to move the source folder into.
	//
	// For example - /path/of/destination/folder/
	DestinationPath string `json:"destinationPath"`
}

//
// RESPONSES
//

type MoveFolderResponse struct {
	JobID string `json:"jobId"`
}

//
// METHODS
//

// MoveFolder will move one folder into another.
func (s *MediaService) MoveFolder(ctx context.Context, r *MoveFolderRequest) (*MoveFolderResponse, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(r)
	if err != nil {
		return nil, err
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/bulkJobs/moveFolder", b, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Submit the request
	res := new(MoveFolderResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
