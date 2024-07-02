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

type RemoveTagsRequest struct {
	// FileIDs is the list of unique ID of the uploaded files.
	FileIDs []string `json:"fileIds"`
	// Tags is an array of tags to add on these files.
	Tags []string `json:"tags"`
}

//
// METHODS
//

// RemoveTags from multiple files in a single request.
func (s *MediaService) RemoveTags(ctx context.Context, r *RemoveTagsRequest) error {
	if r == nil {
		return errors.New("request is empty")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(r)
	if err != nil {
		return err
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/removeTags", b, requestTypeAPI)
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
