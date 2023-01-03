package fileput

type Config struct {
	FileServiceAddr string `environment:"FILE_SERVICE_ADDR" json:"core-file-addr" default:":13007"`
}
