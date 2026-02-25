package services

import (
	"context"
	"fmt"
	"reflect"
)

func Call(target interface{}, args ...interface{}) ([]interface{}, error) {
	return CallForContext(context.Background(), target, args...)
}

func CallForContext(ctx context.Context, target interface{}, args ...interface{}) (results []interface{}, err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Func {
		callResults := invokeFunction(ctx, targetValue, args...)
		results = make([]interface{}, len(callResults))
		for i := range callResults {
			results[i] = callResults[i].Interface()
		}
		return results, nil
	}
	return nil, fmt.Errorf("Only functions can be invoked")
}
