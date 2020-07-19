package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func JSON(ctx *fasthttp.RequestCtx, statusCode int, data interface{}) {
	ctx.SetContentType("application/json")
	ctx.Response.SetStatusCode(statusCode)
	if err := json.NewEncoder(ctx).Encode(data); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusBadRequest)
	}
}

func ERROR(ctx *fasthttp.RequestCtx, statusCode int, err error) {
	ctx.SetContentType("application/json")
	if err != nil {
		JSON(ctx, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(ctx, fasthttp.StatusBadRequest, fasthttp.StatusMessage(fasthttp.StatusBadRequest))
}
