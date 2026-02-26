package services

import (
	"context"
	"fmt"
	"reflect"
)

func Populate(targetPtr interface{}) error {
	return PopulateForContext(context.Background(), targetPtr)
}

func PopulateForContext(
	ctx context.Context,
	targetPtr interface{},
) error {
	return PopulateForContextWithExtras(ctx, targetPtr, make(map[reflect.Type]reflect.Value))
}

func PopulateForContextWithExtras(
	ctx context.Context,
	targetPtr interface{},
	extras map[reflect.Type]reflect.Value,
) error {
	v := reflect.ValueOf(targetPtr)
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("expected pointer to struct, got %T", targetPtr)
	}
	targetValue := v.Elem()
	for i := 0; i < targetValue.NumField(); i++ {
		fieldVal := targetValue.Field(i)
		if !fieldVal.CanSet() {
			continue
		}
		// Only resolve interface fields (service dependencies)
		if fieldVal.Kind() != reflect.Interface {
			continue
		}
		if extra, ok := extras[fieldVal.Type()]; ok {
			fieldVal.Set(extra)
			continue
		}
		if err := resolveServiceFromValue(ctx, fieldVal.Addr()); err != nil {
			return fmt.Errorf("field %s: %w", targetValue.Type().Field(i).Name, err)
		}
	}
	return nil
}
