package imagekit

import (
	"context"
	"errors"
)

//
// METHODS
//

// GetForUploaded gets image exif, pHash and other metadata for uploaded files in Imagekit.io media library using this API.
func (s *MetadataService) GetForUploaded(ctx context.Context, fid string) (*MetadataResponse, error) {
	if fid == "" {
		return nil, errors.New("file id is required")
	}

	// Prepare request
	req, err := s.client.request("GET", "v1/files/"+fid+"/metadata", nil, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Submit the request
	res := new(MetadataResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
