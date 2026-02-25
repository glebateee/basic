package services

import (
	"context"
	"fmt"
	"reflect"
)

func GetService(targetPtr interface{}) error {
	return GetServiceForContext(context.Background(), targetPtr)
}

func GetServiceForContext(ctx context.Context, targetPtr interface{}) error {
	targetValue := reflect.ValueOf(targetPtr)
	if targetValue.Kind() == reflect.Pointer && targetValue.Elem().CanSet() {
		return resolveServiceFromValue(ctx, targetValue)
	}
	return fmt.Errorf("type cannot be used as target")
}
