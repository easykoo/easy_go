package handler

import (
	"middleware"
)

func IndexHandler(ctx *middleware.Context) {
	ctx.HTML(200, "index", ctx)
}
