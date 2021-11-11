package send_email

import _ "embed"

//go:embed password_recovery.html
var passwordRecoveryTmpl string

var passwordRecovery = emailContent{
	subject: `Reset your password`,
	body:    "{{ define content }}" + passwordRecoveryTmpl + "{{ end }}",
}
