package actionresults

import "encoding/json"

type JsonActionResult struct {
	data interface{}
}

func NewJsonAction(data interface{}) ActionResult {
	return &JsonActionResult{data: data}
}

func (j *JsonActionResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(ctx.ResponseWriter).Encode(j.data)
}
