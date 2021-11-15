package send_email

import (
	_ "embed"
	"html/template"
)

//go:embed invite_user.html
var inviteUserTmpl string

var inviteUser = emailContent{
	subject: template.Must(template.New("inviteUserSubject").Parse(`You've been invited by {{ .Sender.Name }} to {{ .AppName }}`)),
	body:    template.Must(template.New("inviteUserBody").Parse(inviteUserTmpl)),
}
