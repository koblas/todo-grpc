package shared_config

import (
	_ "embed"
)

// /
// JSON Configuration that contains some "well known" handles to make sure they're
// not repeated in code for the heck of it

//go:embed config.json
var CONFIG string
