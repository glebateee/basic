package basic

import (
	"github.com/glebateee/basic/internal/pipeline"
	"github.com/glebateee/basic/internal/services"
)

type ServicesComponent struct{}

func (c *ServicesComponent) Init() {}

func (c *ServicesComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	reqCtx := ctx.Request.Context()
	ctx.Request = ctx.Request.WithContext(services.NewServiceContext(reqCtx))
	next(ctx)
}
