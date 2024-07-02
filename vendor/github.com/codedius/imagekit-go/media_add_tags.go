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

type AddTagsRequest struct {
	// FileIDs is the list of unique ID of the uploaded files.
	FileIDs []string `json:"fileIds"`
	// Tags is an array of tags to add on these files.
	Tags []string `json:"tags"`
}

//
// METHODS
//

// AddTags to multiple files in a single request.
func (s *MediaService) AddTags(ctx context.Context, r *AddTagsRequest) error {
	if r == nil {
		return errors.New("request is empty")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(r)
	if err != nil {
		return nil
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/addTags", b, requestTypeAPI)
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
