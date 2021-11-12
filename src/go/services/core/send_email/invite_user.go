package send_email

import _ "embed"

//go:embed invite_user.html
var inviteUserTmpl string

var inviteUser = emailContent{
	subject: `You've been invited by {{ .Sender.Name }} to {{ .AppName }}`,
	body:    "{{ define content }}" + inviteUserTmpl + "{{ end }}",
}
