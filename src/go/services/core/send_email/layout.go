package send_email

import (
	_ "embed"
	"html/template"
)

//go:embed layout.html
var baseLayout string

var _layout = template.Must(template.New("layout").Parse(baseLayout))

// Layout -- exported for testing
func Layout() *template.Template {
	return template.Must(_layout.Clone())
}
