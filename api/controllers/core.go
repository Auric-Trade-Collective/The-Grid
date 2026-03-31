package controllers

import (
	"github.com/Auric-Trade-Collective/The-Grid/api/types"
	noxgo "github.com/Auric-Trade-Collective/nox-go"
)

type CoreController struct{}

func (c *CoreController) Health(resp *noxgo.HttpResponse, req *noxgo.HttpRequest) {
	resp.WriteJson(types.NewApiResponse("OK", []any{}))
}
