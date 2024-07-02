package imagekit

//
// RESPONSES
//

type MetadataResponse struct {
	Density         int                   `json:"density"`
	EXIF            *MetadataResponseEXIF `json:"exif"`
	Format          string                `json:"format"`
	HasColorProfile bool                  `json:"hasColorProfile"`
	HasTransparency bool                  `json:"hasTransparency"`
	Height          int                   `json:"height"`
	PHash           string                `json:"pHash"`
	Quality         int                   `json:"quality"`
	Size            int                   `json:"size"`
	Width           int                   `json:"width"`
}

type MetadataResponseEXIF struct {
	EXIF             *MetadataEXIF             `json:"exif"`
	GPS              *MetadataGPS              `json:"gps"`
	Image            *MetadataImage            `json:"image"`
	Interoperability *MetadataInteroperability `json:"interoperability"`
	Makernote        *MetadataMakernote        `json:"makernote"`
	Thumbnail        *MetadataThumbnail        `json:"thumbnail"`
}

type MetadataEXIF struct {
	ApertureValue            float32 `json:"ApertureValue"`
	ColorSpace               int     `json:"ColorSpace"`
	CreateDate               string  `json:"CreateDate"`
	CustomRendered           int     `json:"CustomRendered"`
	DateTimeOriginal         string  `json:"DateTimeOriginal"`
	ExifImageHeight          int     `json:"ExifImageHeight"`
	ExifImageWidth           int     `json:"ExifImageWidth"`
	ExifVersion              string  `json:"ExifVersion"`
	ExposureCompensation     int     `json:"ExposureCompensation"`
	ExposureMode             int     `json:"ExposureMode"`
	ExposureProgram          int     `json:"ExposureProgram"`
	ExposureTime             float32 `json:"ExposureTime"`
	FNumber                  float32 `json:"FNumber"`
	Flash                    int     `json:"Flash"`
	FlashpixVersion          string  `json:"FlashpixVersion"`
	FocalLength              int     `json:"FocalLength"`
	FocalPlaneResolutionUnit int     `json:"FocalPlaneResolutionUnit"`
	FocalPlaneXResolution    float32 `json:"FocalPlaneXResolution"`
	FocalPlaneYResolution    float32 `json:"FocalPlaneYResolution"`
	ISO                      int     `json:"ISO"`
	InteropOffset            int     `json:"InteropOffset"`
	MeteringMode             int     `json:"MeteringMode"`
	SceneCaptureType         int     `json:"SceneCaptureType"`
	ShutterSpeedValue        float32 `json:"ShutterSpeedValue"`
	SubSecTime               string  `json:"SubSecTime"`
	SubSecTimeDigitized      string  `json:"SubSecTimeDigitized"`
	SubSecTimeOriginal       string  `json:"SubSecTimeOriginal"`
	WhiteBalance             int     `json:"WhiteBalance"`
}

type MetadataGPS struct {
	GPSVersionID []int `json:"GPSVersionID"`
}

type MetadataImage struct {
	EXIFOffset       int    `json:"ExifOffset"`
	GPSInfo          int    `json:"GPSInfo"`
	Make             string `json:"Make"`
	Model            string `json:"Model"`
	ModifyDate       string `json:"ModifyDate"`
	Orientation      int    `json:"Orientation"`
	ResolutionUnit   int    `json:"ResolutionUnit"`
	Software         string `json:"Software"`
	XResolution      int    `json:"XResolution"`
	YCbCrPositioning int    `json:"YCbCrPositioning"`
	YResolution      int    `json:"YResolution"`
}

type MetadataInteroperability struct {
	InteropIndex   string `json:"InteropIndex"`
	InteropVersion string `json:"InteropVersion"`
}

type MetadataThumbnail struct {
	Compression     int `json:"Compression"`
	ResolutionUnit  int `json:"ResolutionUnit"`
	ThumbnailLength int `json:"ThumbnailLength"`
	ThumbnailOffset int `json:"ThumbnailOffset"`
	XResolution     int `json:"XResolution"`
	YResolution     int `json:"YResolution"`
}

type MetadataMakernote struct {
}

//
// SERVICES
//

// MetadataService handles communication with the metadata related methods of the ImageKit API.
type MetadataService service
