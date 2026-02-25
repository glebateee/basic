package services

import (
	"context"
	"reflect"
)

const ServiceKey = "services"

type serviceMap map[reflect.Type]reflect.Value

func NewServiceContext(ctx context.Context) context.Context {
	if ctx.Value(ServiceKey) == nil {
		return context.WithValue(ctx, ServiceKey, make(serviceMap))
	}
	return ctx
}
