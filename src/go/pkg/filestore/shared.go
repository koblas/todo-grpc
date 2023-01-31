package filestore

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

func buildObjectKey(params *FilePutParams) string {
	path := strings.Replace(strings.Join([]string{
		params.Prefix,
		params.FileType,
		params.UserId,
		ulid.Make().String() + params.Suffix,
	}, "/"), "//", "/", -1)

	return strings.TrimPrefix(path, "/")
}
