package imagekit

import (
	"bytes"
	"context"
	"errors"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
)

//
// METHODS
//

// ServerUpload uploads file to ImageKit.io.
func (s *UploadService) ServerUpload(ctx context.Context, r *UploadRequest) (*UploadResponse, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}
	if r.FileName == "" {
		return nil, errors.New("file name is required")
	}

	// Prepare body
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Compose all fields
	err := w.WriteField("fileName", r.FileName)
	if err != nil {
		return nil, err
	}
	err = w.WriteField("useUniqueFileName", strconv.FormatBool(r.UseUniqueFileName))
	if err != nil {
		return nil, err
	}
	err = w.WriteField("tags", strings.Join(r.Tags, ","))
	if err != nil {
		return nil, err
	}
	err = w.WriteField("folder", r.Folder)
	if err != nil {
		return nil, err
	}
	err = w.WriteField("isPrivateFile", strconv.FormatBool(r.IsPrivateFile))
	if err != nil {
		return nil, err
	}
	err = w.WriteField("customCoordinates", r.CustomCoordinates)
	if err != nil {
		return nil, err
	}
	err = w.WriteField("responseFields", strings.Join(r.ResponseFields, ","))
	if err != nil {
		return nil, err
	}

	// Add file content to the request form
	if fileBytes, ok := r.File.([]byte); ok {
		file, err := w.CreateFormFile("file", r.FileName)
		if err != nil {
			return nil, err
		}
		_, err = file.Write(fileBytes)
		if err != nil {
			return nil, err
		}
	} else if fileURL, ok := r.File.(*url.URL); ok {
		err = w.WriteField("file", fileURL.String())
		if err != nil {
			return nil, err
		}
	} else if fileURL, ok := r.File.(url.URL); ok {
		err = w.WriteField("file", fileURL.String())
		if err != nil {
			return nil, err
		}
	} else if fileBase64, ok := r.File.(string); ok {
		err = w.WriteField("file", fileBase64)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("file type cannot be defined (only []byte, base64 string and URL allowed)")
	}

	err = w.Close()
	if err != nil {
		return nil, err
	}

	// Prepare request
	req, err := s.client.request("POST", "api/v1/files/upload", &b, requestTypeUpload)
	if err != nil {
		return nil, err
	}

	// Set necessary headers
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res := new(UploadResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, err
}
