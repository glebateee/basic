package actionresults

import "net/http"

type RedirectActionResult struct {
	url string
}

func NewRedirectAction(url string) ActionResult {
	return &RedirectActionResult{url: url}
}

func (r *RedirectActionResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Location", r.url)
	ctx.ResponseWriter.WriteHeader(http.StatusSeeOther)
	return nil
}
