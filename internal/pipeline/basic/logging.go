package basic

import (
	"net/http"

	"github.com/glebateee/basic/internal/logging"
	"github.com/glebateee/basic/internal/pipeline"
)

var emptyStatus = 0

type LoggingResponseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (rw *LoggingResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *LoggingResponseWriter) Write(b []byte) (int, error) {
	if rw.statusCode == emptyStatus {
		rw.statusCode = http.StatusOK
	}
	return rw.ResponseWriter.Write(b)
}

type LoggingComponent struct{}

func (lc *LoggingComponent) Init() {}

func (lc *LoggingComponent) ImplementsProcessRequestWithServices() {}

func (lc *LoggingComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	logger logging.Logger,
) {

	wrapper := LoggingResponseWriter{
		statusCode:     emptyStatus,
		ResponseWriter: ctx.ResponseWriter,
	}

	ctx.ResponseWriter = &wrapper
	logger.Infof("REQ --- %v - %v", ctx.Request.Method, ctx.Request.URL)
	next(ctx)
	logger.Infof("RSP  %v  %v", wrapper.statusCode, ctx.Request.URL)
}
