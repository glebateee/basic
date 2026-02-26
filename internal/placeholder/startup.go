package placeholder

import (
	"sync"

	"github.com/glebateee/basic/internal/http"
	"github.com/glebateee/basic/internal/pipeline"
	"github.com/glebateee/basic/internal/pipeline/basic"
	"github.com/glebateee/basic/internal/services"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&SimpleMessageComponent{},
	)
}

func Start() {
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
