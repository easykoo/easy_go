package handler

import (
	"github.com/go-martini/martini"

	. "common"
	"middleware"
	"model"
)

func IndexHandler(ctx *middleware.Context) {
	ctx.HTML(200, "index", ctx)
}

func LangHandler(ctx *middleware.Context, params martini.Params) {
	lang := params["lang"]
	ctx.SessionSet("Lang", lang)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func Contact(ctx *middleware.Context, feedback model.Feedback) {
	err := feedback.Insert()
	PanicIf(err)
	ctx.Set("success", true)
	ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.send.success"))
	ctx.JSON(200, ctx.Response)
}
