package filestore

import (
	"context"
	"io"
)

// type FileUploadUrlParams struct {
// 	UserId string
// 	Type   string
// }
// type FileUploadUrlResponse struct {
// 	Url string
// }

type FilePutParams struct {
	Bucket      string
	Prefix      string
	UserId      string
	FileType    string
	Suffix      string
	ContentType string
}
type FilePutResponse struct {
	Url string
}

type FileGetParams struct {
	Path string
}
type FileGetResponse struct {
	Data []byte
}

type Filestore interface {
	UploadUrl(ctx context.Context, params *FilePutParams) (*FilePutResponse, error)
	PutFile(ctx context.Context, params *FilePutParams, reader io.Reader) (*FilePutResponse, error)
	GetFile(ctx context.Context, params *FileGetParams) (io.Reader, error)
}
