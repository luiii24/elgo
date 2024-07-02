package imagekit

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"time"
)

//
// REQUESTS
//

type UpdateFileDetailsRequest struct {
	// Tags associated with the file.
	Tags []string `json:"tags"`
	// CustomCoordinates define an important area in the image. This is only relevant for image type files.
	CustomCoordinates string `json:"customCoordinates"`
}

//
// RESPONSES
//

type UpdateFileDetailsResponse struct {
	// FileID is the unique ID of the uploaded file.
	FileID string `json:"fileId"`
	// Type of item. It can be either file or imageFolder.
	Type string `json:"type"`
	// Name of the file or imageFolder.
	Name string `json:"name"`
	// FilePath of the file. In the case of an image, you can use this path to construct different transform.
	FilePath string `json:"filePath"`
	// Tags is array of tags associated with the image.
	Tags []string `json:"tags"`
	// IsPrivateFile is the file marked as private. It can be either "true" or "false".
	IsPrivateFile bool `json:"isPrivateFile"`
	// CustomCoordinates is the value of custom coordinates associated with the image in format "x,y,width,height".
	CustomCoordinates string `json:"customCoordinates"`
	// URL of the file.
	URL string `json:"url"`
	// Thumbnail is a small thumbnail URL in case of an image.
	Thumbnail string `json:"thumbnail"`
	// FileType of the file, it could be either image or non-image.
	FileType string `json:"fileType"`
	// MIME Type of the file.
	MIME string `json:"mime"`
	// Height of the uploaded image file.
	//
	// Only applicable when file type is image.
	Height int `json:"height"`
	// Width of the uploaded image file.
	//
	// Only applicable when file type is image.
	Width int `json:"width"`
	// Size of the uploaded file in bytes.
	Size int `json:"size"`
	// HasAlpha is whether the image has an alpha component or not.
	HasAlpha bool `json:"hasAlpha"`
	// The date and time when the file was first uploaded.
	//
	// The format is YYYY-MM-DDTHH:mm:ss.sssZ
	CreatedAt time.Time `json:"created_at"`
}

//
// METHODS
//

// UpdateFileDetails such as tags, customCoordinates, and isPrivate properties using get file detail API.
func (s *MediaService) UpdateFileDetails(ctx context.Context, fid string, r *UpdateFileDetailsRequest) (*UpdateFileDetailsResponse, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}
	if fid == "" {
		return nil, errors.New("file id is emoty")
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(r)
	if err != nil {
		return nil, err
	}

	// Prepare request
	req, err := s.client.request("PATCH", "v1/files/"+fid+"/details", b, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Submit the request
	res := new(UpdateFileDetailsResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
