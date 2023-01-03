package user

type SsmConfig struct {
	UserServiceAddr string `environment:"USER_SERVICE_ADDR" json:"core-user-addr" default:":13005"`
	JwtSecret       string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
}
