# imagekit-go
A Go client library for accessing the [ImageKit.io API](https://docs.imagekit.io).

## Installation
Use the following command to download this module:
```
go get github.com/codedius/imagekit-go
```

## Usage
```go
import "github.com/codedius/imagekit-go"
```
Construct a new API client, then use to access the ImageKit.io API. For example:
```go
opts := imagekit.Options{
    PublicKey:  "Your API Private Key",
    PrivateKey: "Your API Public Key",
}

ik, err := imagekit.NewClient(&opts)
if err != nil {
    // error handling
}
```
Upload image to your ImageKit.io Media Library:
```go
ur := imagekit.UploadRequest{
    File:              file, // []byte OR *url.URL OR url.URL OR base64 string
    FileName:          "file name",
    UseUniqueFileName: false,
    Tags:              []string{},
    Folder:            "/",
    IsPrivateFile:     false,
    CustomCoordinates: "",
    ResponseFields:    nil,
}

ctx := context.Background()

upr, err := ik.Upload.ServerUpload(ctx, &ur)
if err != nil {
    // error handling
}
```
Other methods are pretty straightforward and doesn't need extra explanations. For more info please read library's [documentation](https://pkg.go.dev/github.com/codedius/imagekit-go?tab=doc) from go.dev and ImageKit.io API [documentation](https://docs.imagekit.io).
