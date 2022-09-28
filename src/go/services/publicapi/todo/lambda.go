package todo

type SsmConfig struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
}
