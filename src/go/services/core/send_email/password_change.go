package send_email

import _ "embed"

//go:embed password_change.html
var passwordChangeTmpl string

var passwordChange = emailContent{
	subject: `Your password was changed`,
	body:    `{{ define "content" }}` + passwordChangeTmpl + "{{ end }}",
}
