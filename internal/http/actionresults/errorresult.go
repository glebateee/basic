package actionresults

type ErrorActionResult struct {
	error
}

func NewErrorAction(err error) ActionResult {
	return &ErrorActionResult{err}
}

func (e *ErrorActionResult) Execute(*ActionContext) error {
	return e.error
}
