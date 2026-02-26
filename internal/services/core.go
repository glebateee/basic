package services

import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle
}

var services = make(map[reflect.Type]BindingMap)

func addService(life lifecycle, factoryFunc interface{}) error {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		returnType := factoryFuncType.Out(0)
		services[returnType] = BindingMap{
			factoryFunc: reflect.ValueOf(factoryFunc),
			lifecycle:   life,
		}
		return nil
	}
	return fmt.Errorf("Type cannot be used as service: %s", factoryFuncType.String())
}

func resolveServiceFromValue(ctx context.Context, targetPtr reflect.Value) error {
	targetType := targetPtr.Elem().Type()
	fmt.Printf("Resolving service for type: %v\n", targetType) // debug
	if targetType == reflect.TypeOf(context.Context(nil)) {
		targetPtr.Elem().Set(reflect.ValueOf(ctx))
		return nil
	}
	if binding, found := services[targetType]; found {
		fmt.Printf("Found binding for %v\n", targetType) // debug
		if binding.lifecycle == Scoped {
			return resolveScopedService(ctx, targetPtr, binding)
		}
		targetPtr.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
		return nil
	}
	return fmt.Errorf("cannot find service %s", targetType.String())
}

func resolveScopedService(ctx context.Context, target reflect.Value, binding BindingMap) error {
	sMap, ok := ctx.Value(ServiceKey).(serviceMap)
	if !ok {
		target.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
	}
	serviceVal, ok := sMap[target.Type()]
	if !ok {
		serviceVal = invokeFunction(ctx, binding.factoryFunc)[0]
		sMap[target.Type()] = serviceVal
	}
	target.Elem().Set(serviceVal)
	return nil
}

func resolveFunctionArguments(ctx context.Context, target reflect.Value, args ...interface{}) []reflect.Value {
	funcParams := make([]reflect.Value, target.Type().NumIn())
	i := 0
	for _, p := range args {
		funcParams[i] = reflect.ValueOf(p)
		i++
	}
	for ; i < len(funcParams); i++ {
		pType := target.Type().In(i)
		pValue := reflect.New(pType)
		err := resolveServiceFromValue(ctx, pValue)
		if err != nil {
			panic(err)
		}
		funcParams[i] = pValue.Elem()
	}
	return funcParams
}

func invokeFunction(ctx context.Context, target reflect.Value, args ...interface{}) []reflect.Value {
	functionParams := resolveFunctionArguments(ctx, target, args...)
	return target.Call(functionParams)
}
