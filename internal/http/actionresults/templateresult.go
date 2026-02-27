package actionresults

import "github.com/glebateee/basic/internal/templates"

type TemplateActionResult struct {
	templateName string
	data         interface{}
	templates.TemplateExecutor
}

func NewTemplateAction(name string, data interface{}) ActionResult {
	return &TemplateActionResult{templateName: name, data: data}
}

func (t *TemplateActionResult) Execute(ctx *ActionContext) error {
	return t.TemplateExecutor.ExecTemplate(ctx.ResponseWriter, t.templateName, t.data)
}
