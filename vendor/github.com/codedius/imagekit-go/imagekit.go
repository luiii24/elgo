// Package imagekit provides Go client to work with Imagekit.io image processing service API.
package imagekit

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//
// CONSTS
//

const (
	baseURLAPI       = "https://api.imagekit.io/"
	baseURLUpload    = "https://upload.imagekit.io/"
	libraryVersion   = "1.1.2"
	libraryUserAgent = "imagekit-go/" + libraryVersion
)

type requestType string

const (
	requestTypeAPI    requestType = "API"
	requestTypeUpload requestType = "Upload"
)

//
// OPTIONS
//

// Options contains all necessary data to make requests to the API.
type Options struct {
	PublicKey  string
	PrivateKey string
}

//
// SERVICE
//

type service struct {
	client *Client
}

//
// CLIENT
//

// Client manages communication with the API.
type Client struct {
	options       *Options
	client        *http.Client
	apiBaseURL    *url.URL
	uploadBaseURL *url.URL

	// Upload API.
	Upload *UploadService
	// Media API.
	Media *MediaService
	// Metadata API.
	Metadata *MetadataService
}

// NewClient returns a new API client.
func NewClient(opts *Options) (*Client, error) {
	if opts == nil {
		return nil, errors.New("options are empty")
	}
	if opts.PublicKey == "" {
		return nil, errors.New("public key is empty")
	}
	if opts.PrivateKey == "" {
		return nil, errors.New("private key is empty")
	}

	httpClient := http.DefaultClient

	abu, _ := url.Parse(baseURLAPI)
	ubu, _ := url.Parse(baseURLUpload)

	c := &Client{
		options:       opts,
		client:        httpClient,
		apiBaseURL:    abu,
		uploadBaseURL: ubu,
	}

	c.Upload = &UploadService{client: c}
	c.Media = &MediaService{client: c}
	c.Metadata = &MetadataService{client: c}

	return c, nil
}

func (c *Client) request(method, path string, body io.Reader, t requestType) (*http.Request, error) {
	p, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := new(url.URL)

	switch t {
	case requestTypeAPI:
		u = c.apiBaseURL.ResolveReference(p)
	case requestTypeUpload:
		u = c.uploadBaseURL.ResolveReference(p)
	default:
		return nil, errors.New("request type is not defined")
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	key := base64.StdEncoding.EncodeToString([]byte(c.options.PrivateKey + ":"))

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", key))
	req.Header.Set("User-Agent", libraryUserAgent)

	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return err
	}
	defer shut(resp.Body)

	err = checkResponse(resp)
	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	if err == io.EOF {
		return nil // ignore EOF errors caused by empty response body
	}

	return err
}

//
// ERRORS
//

// ErrorResponse reports error caused by an API request.
type ErrorResponse struct {
	Response       *http.Response
	Message        string   `json:"message"`
	Help           string   `json:"help"`
	MissingFileIDs []string `json:"missingFileIds,omitempty"`
}

// Errors provides string error of error repsponse r.
func (r *ErrorResponse) Error() string {
	message := fmt.Sprintf("%v %v: %d", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode)

	if r.Message != "" {
		message = fmt.Sprintf("%s %s", message, r.Message)
	}
	if r.Help != "" {
		message = fmt.Sprintf("%s (%s)", message, r.Help)
	}
	if len(r.MissingFileIDs) > 0 {
		message = fmt.Sprintf("%s (missing IDs: %v)", message, r.MissingFileIDs)
	}

	return message
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	errResp := &ErrorResponse{Response: r}

	if data != nil {
		err = json.Unmarshal(data, errResp)
		if err != nil {
			return err
		}
	}

	return errResp
}

//
// UTILS
//

func shut(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
