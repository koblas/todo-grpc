package file

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrorLookupNotFound = errors.New("not found")

const (
	MEMORY_NEW = iota
	MEMORY_UPLOADED
)

type FileByteStore map[string][]byte

type memoryStore struct {
	fileInfo map[string]*FileInfo
	files    FileByteStore
	prefix   string
}

func NewFileMemoryStore(prefix string) FileStore {
	return &memoryStore{
		prefix:   prefix,
		fileInfo: map[string]*FileInfo{},
		files:    FileByteStore{},
	}
}

func (store *memoryStore) CreateUploadUrl(ctx context.Context, userId, fileType string) (string, error) {
	path := store.prefix + uuid.NewString()
	entry := FileInfo{
		Id:          uuid.NewString(),
		UserId:      userId,
		FileType:    fileType,
		Url:         path,
		InternalUrl: "corefile:" + path,
		expires:     time.Now().Add(time.Duration(5) * time.Minute),
		status:      MEMORY_NEW,
	}

	store.fileInfo[entry.Url] = &entry
	return entry.Url, nil
}

func (store *memoryStore) LookupUploadUrl(ctx context.Context, url string) (*FileInfo, error) {
	entry, found := store.fileInfo[url]
	if !found {
		return nil, ErrorLookupNotFound

	}

	return entry, nil
}

func (store *memoryStore) StoreFile(ctx context.Context, url string, bytes []byte) (string, *FileInfo, error) {
	entry, found := store.fileInfo[url]
	if !found {
		return "", nil, ErrorLookupNotFound
	}
	if entry.status != MEMORY_NEW {
		return "", nil, ErrorLookupNotFound
	}
	store.fileInfo[url].status = MEMORY_UPLOADED

	store.files[url] = bytes

	return url, entry, nil
}

func (store *memoryStore) GetFile(ctx context.Context, path string) ([]byte, error) {
	data, found := store.files[path]
	if !found {
		return nil, ErrorLookupNotFound
	}
	return data, nil
}
