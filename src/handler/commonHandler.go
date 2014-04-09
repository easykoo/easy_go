package handler

import (
	"github.com/go-martini/martini"

	"middleware"
	. "common"
)

func IndexHandler(ctx *middleware.Context) {
	ctx.HTML(200, "index", ctx)
}

func LangHandler(ctx *middleware.Context, params martini.Params) {
	lang := params["lang"]
	Log.Debug("before: ", ctx.Session["Lang"])
	ctx.SessionSet("Lang", lang)
	Log.Debug("after: ", ctx.Session["Lang"])
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}
