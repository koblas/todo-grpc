package file

type SsmConfig struct {
	FileServiceAddr string `environment:"FILE_SERVICE_ADDR" json:"core-file-addr" default:":13007"`
	JwtSecret       string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
}
