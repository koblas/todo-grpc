package send_email

import (
	_ "embed"
	"html/template"
)

//go:embed register_user.html
var registerUserTmpl string

var registerUser = emailContent{
	subject: template.Must(template.New("registerUserSubject").Parse(`Confirm your email address`)),
	body:    template.Must(template.New("registerUserBody").Parse(registerUserTmpl)),
}
