package templates

import (
	"html/template"
	"strings"
)

func insertBodyWrapper(body *strings.Builder) func() template.HTML {
	return func() template.HTML {
		return template.HTML(body.String())
	}
}
func setLayoutWrapper(val *string) func(string) string {
	return func(name string) string {
		*val = name
		return ""
	}
}
