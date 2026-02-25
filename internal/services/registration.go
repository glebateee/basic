package services

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

func AddTransient(factoryFunc interface{}) error {
	return addService(Transient, factoryFunc)
}

func AddScoped(factoryFunc interface{}) error {
	return addService(Scoped, factoryFunc)
}
func AddSingleton(factoryFunc interface{}) error {
	factoryFuncValue := reflect.ValueOf(factoryFunc)
	if factoryFuncValue.Kind() == reflect.Func && factoryFuncValue.Type().NumOut() == 1 {
		var results []reflect.Value
		once := sync.Once{}
		wrapper := reflect.MakeFunc(factoryFuncValue.Type(),
			func([]reflect.Value) []reflect.Value {
				once.Do(func() {
					results = invokeFunction(context.Background(), factoryFuncValue)
				})
				return results
			})
		return addService(Singleton, wrapper.Interface())
	}
	return fmt.Errorf("not a factory function")
}
