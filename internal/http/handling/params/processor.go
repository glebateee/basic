package params

import (
	"net/http"
	"reflect"
)

func GetParametersFromRequest(
	request *http.Request,
	handlerMethod reflect.Method,
	urlVals []string,
) ([]reflect.Value, error) {
	handlerMethodType := handlerMethod.Type
	if handlerMethodType.NumIn() == 1 {
		return []reflect.Value{}, nil
	}
	var err error
	//params := make([]reflect.Value, handlerMethodType.NumIn()-1)
	if handlerMethodType.NumIn() == 2 && handlerMethodType.In(1).Kind() == reflect.Struct {
		structVal := reflect.New(handlerMethodType.In(1))
		err = request.ParseForm()
		if err == nil {
			if getContentType(request) == "application/json" {
				err = populateStructFromJSON(structVal, request.Body)
			} else {
				err = populateStructFromForm(structVal, request.Form)
			}
			return []reflect.Value{structVal.Elem()}, err
		}
	}
	return getParametersFromURLValues(handlerMethodType, urlVals)
}

func getContentType(request *http.Request) string {
	headerSlice := request.Header["Content-Type"]
	if len(headerSlice) > 0 {
		return headerSlice[0]
	}
	return ""
}
