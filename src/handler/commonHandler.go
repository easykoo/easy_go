package handler

import (
	"github.com/go-martini/martini"

	. "common"
	"middleware"
	"model"
)

func Index(ctx *middleware.Context) {
	ctx.HTML(200, "index", ctx)
}

func About(ctx *middleware.Context) {
	ctx.HTML(200, "about", ctx)
}

func ContactHandler(ctx *middleware.Context, feedback model.Feedback) {
	switch ctx.R.Method {
	case "POST":
		err := feedback.Insert()
		PanicIf(err)
		ctx.Set("success", true)
		ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.send.success"))
		ctx.JSON(200, ctx.Response)
	default:
		ctx.HTML(200, "contact", ctx)
	}
}

func LangHandler(ctx *middleware.Context, params martini.Params) {
	lang := params["lang"]
	ctx.SessionSet("Lang", lang)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}
