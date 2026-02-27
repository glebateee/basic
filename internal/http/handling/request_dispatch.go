package handling

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/glebateee/basic/internal/http/actionresults"
	"github.com/glebateee/basic/internal/http/handling/params"
	"github.com/glebateee/basic/internal/pipeline"
	"github.com/glebateee/basic/internal/services"
)

type RouterComponent struct {
	routes []Route
}

func NewRouter(handlers ...HandlerEntry) *RouterComponent {
	return &RouterComponent{generateRoutes(handlers...)}
}

func (rc *RouterComponent) Init() {}

func (rc *RouterComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	for _, route := range rc.routes {
		if strings.EqualFold(route.httpMethod, ctx.Request.Method) {
			matches := route.expression.FindAllStringSubmatch(ctx.URL.Path, -1)
			fmt.Println(matches)
			if len(matches) > 0 {
				var rawParams []string
				if len(matches[0]) > 0 {
					rawParams = matches[0][1:]
				}
				if err := rc.invokeHandler(route, rawParams, ctx); err != nil {
					ctx.NewError(err)
				} else {
					next(ctx)
				}
				return
			}
		}
	}
	ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
}

func (rc *RouterComponent) invokeHandler(
	route Route,
	rawParams []string,
	ctx *pipeline.ComponentContext,
) error {
	paramVals, err := params.GetParametersFromRequest(ctx.Request, route.handlerMethod, rawParams)
	fmt.Println(rawParams)
	fmt.Println(paramVals)
	if err != nil {
		return err
	}
	receiver := reflect.New(route.handlerMethod.Type.In(0))
	if err := services.PopulateForContext(ctx.Request.Context(), receiver.Interface()); err != nil {
		return err
	}
	paramVals = append([]reflect.Value{receiver.Elem()}, paramVals...)
	result := route.handlerMethod.Func.Call(paramVals)
	if len(result) > 0 {
		if action, ok := result[0].Interface().(actionresults.ActionResult); ok {
			if err := services.PopulateForContext(ctx.Request.Context(), action); err != nil {
				return err
			}
			return action.Execute(&actionresults.ActionContext{
				Context:        ctx.Request.Context(),
				ResponseWriter: ctx.ResponseWriter,
			})
		}
		io.WriteString(ctx.ResponseWriter, fmt.Sprint(result[0].Interface()))
	}
	return nil
}
