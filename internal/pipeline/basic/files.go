package basic

import (
	"net/http"
	"os"
	"strings"

	"github.com/glebateee/basic/internal/config"
	"github.com/glebateee/basic/internal/pipeline"
)

type StaticFileComponent struct {
	prefix   string
	provider http.Handler
	Cfg      config.Config
}

func (fc *StaticFileComponent) Init() {
	fc.prefix = fc.Cfg.GetStringDefault("files:urlprefix", "/files/")
	if pathToFiles, ok := fc.Cfg.GetString("files:path"); ok {
		if _, err := os.Stat(pathToFiles); err == nil {
			fc.provider = http.StripPrefix(
				fc.prefix,
				http.FileServer(http.Dir(pathToFiles)),
			)
			return
		}
	}
	panic("Cannot load file configuration settings")
}

func (fc *StaticFileComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	if strings.HasPrefix(ctx.Request.URL.Path, fc.prefix) &&
		!strings.EqualFold(ctx.Request.URL.Path, fc.prefix) {
		fc.provider.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
