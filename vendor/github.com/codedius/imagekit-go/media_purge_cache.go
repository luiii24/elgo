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

type PurgeCacheRequest struct {
	// URL is the exact URL of the file to be purged.
	//
	// For example - https://ik.imageki.io/your_imagekit_id/rest-of-the-file-path.jpg.
	URL string `json:"url"`
}

//
// RESPONSES
//

type PurgeCacheResponse struct {
	// RequestID which can be used to get the purge request status.
	RequestID string `json:"requestId"`
}

//
// METHODS
//

// PurgeCache will purge CDN and Imagekit.io internal cache.
func (s *MediaService) PurgeCache(ctx context.Context, r *PurgeCacheRequest) (*PurgeCacheResponse, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}
	if r.URL == "" {
		return nil, errors.New("URL is empty")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(r)
	if err != nil {
		return nil, err
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/purge", b, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Submit the request
	res := new(PurgeCacheResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
