package params

import (
	"errors"
	"reflect"
)

func getParametersFromURLValues(funcType reflect.Type, urlVals []string) ([]reflect.Value, error) {
	var err error
	if len(urlVals)+1 == funcType.NumIn() {
		params := make([]reflect.Value, len(urlVals))
		for i := range len(urlVals) {
			params[i], err = parseValueToType(funcType.In(i+1), urlVals[i])
			if err != nil {
				return nil, err
			}
		}
		return params, nil
	}
	return nil, errors.New("Parameter number mismatch")
}
