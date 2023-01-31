package file

type Config struct {
	UploadBucket  string `environment:"UPLOAD_BUCKET"`
	JwtSecret     string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	MinioEndpoint string `environment:"MINIO_ENDPOINT" default:"s3.amazonaws.com"`
}
