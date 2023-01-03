package file

import (
	"context"
	"time"
)

type FileInfo struct {
	Id          string
	FileType    string
	UserId      string
	Url         string
	InternalUrl string
	// technically implementation dependant
	expires time.Time
	status  int
}

type FileStore interface {
	CreateUploadUrl(ctx context.Context, userId, fileType string) (string, error)
	LookupUploadUrl(ctx context.Context, url string) (*FileInfo, error)
	StoreFile(ctx context.Context, url string, bytes []byte) (string, *FileInfo, error)
	GetFile(ctx context.Context, path string) ([]byte, error)
}
