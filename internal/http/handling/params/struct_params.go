package params

import (
	"encoding/json"
	"io"
	"net/url"
	"reflect"
	"strings"
)

func populateStructFromForm(structValPtr reflect.Value, formVals url.Values) error {
	var err error
	for i := range structValPtr.Elem().Type().NumField() {
		fieldType := structValPtr.Elem().Type().Field(i)
		for key, vals := range formVals {
			if strings.EqualFold(key, fieldType.Name) && len(vals) > 0 {
				fieldVal := structValPtr.Elem().Field(i)
				if fieldVal.CanSet() {
					valToSet, convErr := parseValueToType(fieldVal.Type(), vals[0])
					if err != nil {
						fieldVal.Set(valToSet)
					} else {
						err = convErr
					}
				}

			}
		}
	}
	return err
}

func populateStructFromJSON(structValPtr reflect.Value, reader io.ReadCloser) (err error) {
	return json.NewDecoder(reader).Decode(structValPtr.Interface())
}
