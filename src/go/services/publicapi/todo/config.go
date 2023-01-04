package todo

type Config struct {
	TodoServiceAddr string `environment:"TODO_SERVICE_ADDR" json:"core-todo-addr" default:":13005"`
	JwtSecret       string `ssm:"jwt_secret" environment:"JWT_SECRET" validate:"min=32"`
}
