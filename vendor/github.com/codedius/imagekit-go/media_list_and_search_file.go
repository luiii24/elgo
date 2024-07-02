package imagekit

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	listAndSearchMinLimit = 1
	listAndSearchMaxLimit = 1000
	listAndSearchMinSkip  = 0
)

//
// REQUESTS
//

type ListAndSearchFileRequest struct {
	// Path if you want to limit the search within a specific imageFolder.
	//
	// For example, /sales-banner/ will only search in imageFolder sales-banner.
	Path string
	// FileType to include in result set.
	//
	// Accepts three values:
	// all - include all types of files in result set
	// image - only search in image type files
	// non-image - only search in files which are not image, e.g., JS or CSS or video files.
	// Default value - all.
	FileType string
	// Files matching any of the Tags are included in result response.
	//
	// If no tag is matched, the file is not included in result set.
	Tags []string
	// IncludeFolder in search results or not. By default only files are searched.
	//
	// Accepts true and false. If this is set to true then tags and FileType parameters are ignored.
	IncludeFolder bool
	// Name of the file or imageFolder.
	Name string
	// Limit the maximum number of results to return in response.
	//
	// Minimum value - 1
	// Maximum value - 1000
	// Default value - 1000
	Limit int
	// Skip the number of results before returning results.
	//
	// Minimum value - 0
	// Default value - 0
	Skip int
}

//
// RESPONSES
//

type ListAndSearchFileResponse struct {
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

// ListAndSearchFile lists all the uploaded files in your Imagekit.io media library.
func (s *MediaService) ListAndSearchFile(ctx context.Context, r *ListAndSearchFileRequest) (*[]ListAndSearchFileResponse, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}

	u, err := url.Parse("v1/files")
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}
	if r.Path != "" {
		parameters.Add("path", r.Path)
	}
	if r.FileType != "" {
		parameters.Add("fileType", r.FileType)
	}
	if len(r.Tags) > 0 {
		parameters.Add("tags", strings.Join(r.Tags, ","))
	}
	if r.Name != "" {
		parameters.Add("name", r.Name)
	}
	if r.Limit < listAndSearchMinLimit {
		parameters.Add("limit", strconv.Itoa(listAndSearchMinLimit))
	} else if r.Limit > listAndSearchMaxLimit {
		parameters.Add("limit", strconv.Itoa(listAndSearchMaxLimit))
	} else {
		parameters.Add("limit", strconv.Itoa(r.Limit))
	}
	if r.Skip < listAndSearchMinSkip {
		parameters.Add("skip", strconv.Itoa(listAndSearchMinSkip))
	} else {
		parameters.Add("skip", strconv.Itoa(r.Skip))
	}
	parameters.Add("includeFolder", strconv.FormatBool(r.IncludeFolder))

	u.RawQuery = parameters.Encode()

	// Prepare request
	req, err := s.client.request("GET", u.String(), nil, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Submit the request
	res := new([]ListAndSearchFileResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
