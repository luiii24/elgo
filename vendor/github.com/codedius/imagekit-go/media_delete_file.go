package imagekit

import (
	"context"
	"errors"
)

//
// METHODS
//

// DeleteFile deletes file with id and all its transform.
func (s *MediaService) DeleteFile(ctx context.Context, fid string) error {
	if fid == "" {
		return errors.New("file id is empty")
	}

	// Prepare request
	req, err := s.client.request("DELETE", "v1/files/"+fid, nil, requestTypeAPI)
	if err != nil {
		return err
	}

	err = s.client.do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
