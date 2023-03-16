package filestore

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

func buildObjectKey(params *FilePutParams) (string, string) {
	id := ulid.Make().String()
	path := strings.Replace(strings.Join([]string{
		params.Prefix,
		params.FileType,
		params.UserId,
		id + params.Suffix,
	}, "/"), "//", "/", -1)

	return strings.TrimPrefix(path, "/"), id
}
