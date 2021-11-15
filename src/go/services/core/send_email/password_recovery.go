package send_email

import (
	_ "embed"
	"html/template"
)

//go:embed password_recovery.html
var passwordRecoveryTmpl string

var passwordRecovery = emailContent{
	subject: template.Must(template.New("passwordRecoverySubject").Parse(`Reset your password`)),
	body:    template.Must(template.New("passwordRecoveryBody").Parse(passwordRecoveryTmpl)),
}
