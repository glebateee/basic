package placeholder

import (
	"github.com/glebateee/basic/internal/config"
	"github.com/glebateee/basic/internal/pipeline"
	"github.com/glebateee/basic/internal/templates"
)

type SimpleMessageComponent struct {
	Message string
	config.Config
}

func (mc *SimpleMessageComponent) Init() {
	mc.Message = mc.Config.GetStringDefault("main:message", "default msg")
}

func (mc *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

func (mc *SimpleMessageComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	executor templates.TemplateExecutor,
) {
	if err := executor.ExecTemplate(ctx.ResponseWriter, "simple_message.html", mc.Message); err != nil {
		ctx.NewError(err)
		return
	}
	next(ctx)
}
