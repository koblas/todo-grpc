package send_email

import (
	_ "embed"
	"html/template"
)

//go:embed password_change.html
var passwordChangeTmpl string

var passwordChange = emailContent{
	subject: template.Must(template.New("passwordChangeSubject").Parse(`Your password was changed`)),
	body:    template.Must(template.New("passwordChangeBody").Parse(passwordChangeTmpl)),
}
