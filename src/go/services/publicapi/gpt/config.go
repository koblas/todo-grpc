package gpt

type Config struct {
	JwtSecret string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	GptApiKey string `ssm:"gpt_api_key" environment:"GPT_API_KEY" validate:"min=2"`
}
