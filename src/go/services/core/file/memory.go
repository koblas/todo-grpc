package file

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

var ErrorLookupNotFound = errors.New("not found")
var ErrorSignatureMismatch = errors.New("signature mismatch")
var SIGNATURE_PARAM = "sig"
var MEMORY_SCHEME = "corefile:"

const (
	MEMORY_NEW = iota
	MEMORY_UPLOADED
)

type FileByteStore map[string][]byte

type memoryStore struct {
	fileInfo map[string]*FileInfo
	files    FileByteStore
	prefix   string
	secret   string
}

var _ FileStore = (*memoryStore)(nil)

func NewFileMemoryStore(prefix string) FileStore {
	return &memoryStore{
		secret:   uuid.NewString(),
		prefix:   prefix,
		fileInfo: map[string]*FileInfo{},
		files:    FileByteStore{},
	}
}

func (store *memoryStore) computeSig(path string) string {
	hasher := sha1.New()
	hasher.Write([]byte(store.secret))
	hasher.Write([]byte(path))
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}

func (store *memoryStore) CreateUploadUrl(ctx context.Context, userId, fileType string) (string, error) {
	path := store.prefix + uuid.NewString()
	entry := FileInfo{
		Id:          uuid.NewString(),
		UserId:      userId,
		FileType:    fileType,
		Url:         path,
		InternalUrl: MEMORY_SCHEME + path,
		expires:     time.Now().Add(time.Duration(5) * time.Minute),
		status:      MEMORY_NEW,
	}

	store.fileInfo[entry.Url] = &entry
	query := "?" + SIGNATURE_PARAM + "=" + store.computeSig(entry.Url)
	return entry.Url + query, nil
}

func (store *memoryStore) LookupUploadUrl(ctx context.Context, url string) (*FileInfo, error) {
	entry, found := store.fileInfo[url]
	if !found {
		return nil, ErrorLookupNotFound
	}

	return entry, nil
}

func (store *memoryStore) VerifyUploadUrl(ctx context.Context, path, query string) error {
	u, err := url.ParseQuery(query)
	if err != nil {
		return ErrorSignatureMismatch
	}

	if u.Get(SIGNATURE_PARAM) != store.computeSig(path) {
		return ErrorSignatureMismatch
	}

	return nil
}

func (store *memoryStore) StoreFile(ctx context.Context, path string, bytes []byte) (string, error) {
	store.files[path] = bytes

	return MEMORY_SCHEME + path, nil
}

func (store *memoryStore) GetFile(ctx context.Context, path string) ([]byte, error) {
	data, found := store.files[strings.TrimPrefix(path, MEMORY_SCHEME)]
	if !found {
		return nil, ErrorLookupNotFound
	}
	return data, nil
}
