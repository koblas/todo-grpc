package auth

type SsmConfig struct {
	RedisAddr            string `ssm:"redis_addr" environment:"REDIS_ADDR"`
	JwtSecret            string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
	UserServiceAddr      string `environment:"USER_SERVICE_ADDR" json:"core-user-addr" default:":13001"`
	OauthUserServiceAddr string `environment:"OAUTH_USER_SERVICE_ADDR" json:"core-oauth-user-addr" default:":13002"`
}
