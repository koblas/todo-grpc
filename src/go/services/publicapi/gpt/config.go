package gpt

type Config struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
}
