package imagekit

import (
	"context"
	"errors"
)

//
// RESPONSES
//

type PurgeCacheStatusResponse struct {
	// Status is the current status of a submitted purge request.
	//
	// It can be either:
	// Pending - The request has been successfully submitted, and purging is in progress.
	// Complete - The purge request has been successfully completed. And now you should get a fresh object. Check the Age header in response to confirm this.
	Status string `json:"status"`
}

//
// METHODS
//

// PurgeCacheStatus gets the status of submitted purge request.
func (s *MediaService) PurgeCacheStatus(ctx context.Context, rid string) (*PurgeCacheStatusResponse, error) {
	if rid == "" {
		return nil, errors.New("request id is empty")
	}

	// Prepare request
	req, err := s.client.request("GET", "v1/files/purge/"+rid, nil, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Submit the request
	res := new(PurgeCacheStatusResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
