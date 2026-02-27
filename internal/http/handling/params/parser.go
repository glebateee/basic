package params

import (
	"fmt"
	"reflect"
	"strconv"
)

func parseValueToType(target reflect.Type, val string) (reflect.Value, error) {
	switch target.Kind() {
	case reflect.Int:
		iVal, err := strconv.Atoi(val)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(iVal), nil
	case reflect.Float64:
		fVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(fVal), nil

	case reflect.Bool:
		bVal, err := strconv.ParseBool(val)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(bVal), nil
	case reflect.String:
		return reflect.ValueOf(val), nil
	}
	return reflect.Value{}, fmt.Errorf("Cannot use type %v as handler method parameter", target.Name())
}
