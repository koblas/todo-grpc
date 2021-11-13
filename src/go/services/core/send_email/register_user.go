package send_email

import _ "embed"

//go:embed register_user.html
var registerUserTmpl string

var registerUser = emailContent{
	subject: `Confirm your email address`,
	body:    `{{ define "content" }}` + registerUserTmpl + "{{ end }}",
}
