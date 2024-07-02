package imagekit

import (
	"context"
	"errors"
	"net/url"
)

//
// METHODS
//

// GetFromRemote gets image exif, pHash and other metadata from Imagekit.io powered remote URL using this API.
func (s *MetadataService) GetFromRemote(ctx context.Context, URL string) (*MetadataResponse, error) {
	if URL == "" {
		return nil, errors.New("URL is required")
	}

	u, err := url.Parse("v1/metadata")
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	parameters.Add("url", URL)

	u.RawQuery = parameters.Encode()

	// Prepare request
	req, err := s.client.request("GET", u.String(), nil, requestTypeAPI)
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
