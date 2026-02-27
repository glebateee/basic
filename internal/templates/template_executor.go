package templates

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data interface{}) error
}

type LayoutTemplateProcessor struct{}

var getTemplates func() *template.Template

func (proc *LayoutTemplateProcessor) ExecTemplate(
	writer io.Writer,
	name string,
	data interface{},
) error {
	var sb strings.Builder
	var layoutName string
	localTemplates := getTemplates()
	localTemplates.Funcs(template.FuncMap{
		"body":   insertBodyWrapper(&sb),
		"layout": setLayoutWrapper(&layoutName),
	})
	if err := localTemplates.ExecuteTemplate(&sb, name, data); err != nil {
		return err
	}
	fmt.Println(sb.String())
	if layoutName != "" {
		return localTemplates.ExecuteTemplate(writer, layoutName, data)
	}
	if _, err := io.WriteString(writer, sb.String()); err != nil {
		return err
	}

	return nil
}
