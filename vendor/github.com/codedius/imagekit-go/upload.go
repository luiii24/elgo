package imagekit

//
// REQUESTS
//

type UploadRequest struct {
	// File to upload.
	File interface{}
	// FileName with which the file has to be uploaded.
	FileName string
	// UseUniqueFileName to whether to use a unique filename for this file or not.
	UseUniqueFileName bool
	// Tags while uploading the file.
	Tags []string
	// Folder path (e.g. /images/imageFolder/) in which the image has to be uploaded. If the imageFolder(s) didn't exist before, a new imageFolder(s) is created.
	Folder string
	// IsPrivateFile to whether to mark the file as private or not. This is only relevant for image type files.
	IsPrivateFile bool
	// CustomCoordinates define an important area in the image. This is only relevant for image type files.
	CustomCoordinates string
	// ResponseFields contains values of the fields that you want ImageKit.io to return in response.
	ResponseFields []string
}

//
// RESPONSES
//

type UploadResponse struct {
	// FileID is unique.
	//
	// Store this fileld in your database, as this will be used to perform update action on this file.
	FileID string `json:"fileId"`
	// Name of the uploaded file.
	Name string `json:"name"`
	// URL of the file.
	URL string `json:"url"`
	// ThumbnailURL is a small thumbnail URL in case of an image.
	ThumbnailURL string `json:"thumbnailUrl"`
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
	// FileType can either be "image" or "non-image".
	FileType string `json:"fileType"`
	// FilePath is the path of the file uploaded.
	//
	// It includes any imageFolder that you specified while uploading.
	FilePath string `json:"filePath"`
	// Tags is array of tags associated with the image.
	Tags []string `json:"tags"`
	// IsPrivateFile is the file marked as private.
	//
	// It can be either "true" or "false".
	IsPrivateFile bool `json:"isPrivateFile"`
	// CustomCoordinates is the value of custom coordinates associated with the image in format "x,y,width,height".
	CustomCoordinates string `json:"customCoordinates"`
	// Metadata of the upload file.
	//
	// Use responseFields property in request to get the metadata returned in response of upload API.
	Metadata interface{} `json:"metadata"`
}

//
// SERVICES
//

// UploadService handles communication with the upload related methods of the ImageKit API.
type UploadService service
