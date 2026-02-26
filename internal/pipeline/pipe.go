package pipeline

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/glebateee/basic/internal/services"
)

type RequestPipeline func(*ComponentContext)

func (pl RequestPipeline) StartPipeline(r *http.Request, w http.ResponseWriter) error {
	ctx := ComponentContext{
		Request:        r,
		ResponseWriter: w,
	}
	pl(&ctx)
	return ctx.error
}

var emptyPipeline RequestPipeline = func(cc *ComponentContext) {}

type MiddlewareComponent interface {
	Init()
	ProcessRequest(ctx *ComponentContext, nf func(*ComponentContext))
}

type ServicesMiddlwareComponent interface {
	Init()
	ImplementsProcessRequestWithServices()
}

func CreatePipeline(components ...interface{}) RequestPipeline {
	ithFunc := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		ithPlus1 := ithFunc
		currComp := components[i]
		if err := services.Populate(currComp); err != nil {
			panic(fmt.Errorf("failed to populate component %T: %w", currComp, err))
		}
		switch currComp := currComp.(type) {
		case MiddlewareComponent:
			ithFunc = createBasicFunction(currComp, ithPlus1)
			currComp.Init()
		case ServicesMiddlwareComponent:
			ithFunc = createServiceDependentFunction(currComp, ithPlus1)
			currComp.Init()
		}

	}
	return ithFunc
}

func createBasicFunction(
	component MiddlewareComponent,
	nextFunc RequestPipeline,
) RequestPipeline {
	return func(cc *ComponentContext) {
		if cc.error == nil {
			component.ProcessRequest(cc, nextFunc)
		}
	}

}
func createServiceDependentFunction(
	component ServicesMiddlwareComponent,
	nextFunc RequestPipeline,
) RequestPipeline {
	method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
	if method.IsValid() {
		return func(ctx *ComponentContext) {
			if ctx.error == nil {
				if _, err := services.CallForContext(
					ctx.Request.Context(),
					method.Interface(),
					ctx,
					nextFunc,
				); err != nil {
					ctx.error = err
				}

			}

		}
	}
	panic("No ProcessRequestWithServices method defined")
}
