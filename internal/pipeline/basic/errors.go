package basic

import (
	"fmt"
	"net/http"

	"github.com/glebateee/basic/internal/logging"
	"github.com/glebateee/basic/internal/pipeline"
)

type ErrorComponent struct{}

func (ec *ErrorComponent) Init() {}

func (ec *ErrorComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	logger logging.Logger,
) {

	defer recoveryPanic(ctx, logger)
	next(ctx)
	if ctx.Error() != nil {
		logger.Debugf("Error: %w", ctx.Error())
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func recoveryPanic(ctx *pipeline.ComponentContext, logger logging.Logger) {
	if arg := recover(); arg != nil {
		logger.Debugf("Error: %s", fmt.Sprint(arg))
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
