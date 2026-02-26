package pipeline

import "net/http"

type ComponentContext struct {
	*http.Request
	http.ResponseWriter
	error
}

func (cc *ComponentContext) Error() error {
	return cc.error
}

func (cc *ComponentContext) NewError(err error) {
	cc.error = err
}
