package handler

import (
	. "common"
	"github.com/go-martini/martini"
	"middleware"
	"model"
)

func PublishBlog(ctx *middleware.Context, blog model.Blog) {
	switch ctx.R.Method {
	case "POST":
		if blog.Title == "" || blog.Content == "" {
			ctx.Set("success", false)
			ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.error.publish.failed"))
		} else {
			err := blog.Insert()
			PanicIf(err)
			ctx.Set("id", blog.Id)
			ctx.Set("success", true)
			ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.publish.success"))
		}
		ctx.JSON(200, ctx.Response)
	default:
		ctx.HTML(200, "blog/publish", ctx)
	}
}

func AllBlog(ctx *middleware.Context) {
	switch ctx.R.Method {
	case "POST":
		blog := new(model.Blog)
		blog.SetPageActive(true)
		blog.SetPageSize(ParseInt(ctx.R.FormValue("iDisplayLength")))
		blog.SetDisplayStart(ParseInt(ctx.R.FormValue("iDisplayStart")))
		columnNum := ctx.R.FormValue("iSortCol_0")
		sortColumn := ctx.R.FormValue("mDataProp_" + columnNum)
		blog.AddSortProperty(sortColumn, ctx.R.FormValue("sSortDir_0"))
		blogArray, total, err := blog.SearchByPage()
		PanicIf(err)
		ctx.Set("aaData", blogArray)
		ctx.Set("iTotalDisplayRecords", total)
		ctx.Set("iTotalRecords", total)
		ctx.JSON(200, ctx.Response)
	default:
		ctx.HTML(200, "blog/allBlog", ctx)
	}
}

func ViewBlog(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	blog := new(model.Blog)
	blog.Id = ParseInt(id)
	err := blog.GetBlog()
	PanicIf(err)
	ctx.Set("blog", blog)
	ctx.HTML(200, "blog/view", ctx)
}
