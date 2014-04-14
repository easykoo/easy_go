package handler

import (
	"github.com/go-martini/martini"
	
	. "common"
	"middleware"
	"model"

	"time"
	"encoding/json"
)

func PublishBlog(ctx *middleware.Context, blog model.Blog) {
	switch ctx.R.Method {
	case "POST":
		if blog.Title == "" || blog.Content == "" {
			ctx.Set("success", false)
			ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.error.publish.failed"))
		} else {
			blog.PublishDate = time.Now()
			blog.State = "PUBLISHED"
			blog.UpdateUser = ctx.SessionGet("SignedUser").(model.User).Username
			if blog.Id == 0 {
				blog.Priority = 5
				blog.CreateUser = ctx.SessionGet("SignedUser").(model.User).Username
				err := blog.Insert()
				PanicIf(err)
			} else {
				err := blog.Update()
				PanicIf(err)
			}
			ctx.Set("blog", blog)
			ctx.Set("success", true)
			ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.publish.success"))
		}
		ctx.JSON(200, ctx.Response)
	default:
		ctx.HTML(200, "blog/publish", ctx)
	}
}

func SaveBlog(ctx *middleware.Context, blog model.Blog) {
	if blog.Title == "" || blog.Content == "" {
		ctx.Set("success", false)
		ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.error.save.failed"))
	} else {
		blog.State = "DRAFT"
		blog.UpdateUser = ctx.SessionGet("SignedUser").(model.User).Username
		if blog.Id == 0 {
			blog.Priority = 5
			blog.CreateUser = ctx.SessionGet("SignedUser").(model.User).Username
			err := blog.Insert()
			PanicIf(err)
		} else {
			err := blog.Update()
			PanicIf(err)
		}
		ctx.Set("blog", blog)
		ctx.Set("success", true)
		ctx.Set("message", Translate(ctx.SessionGet("Lang").(string), "message.save.success"))
	}
	ctx.JSON(200, ctx.Response)
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

func DeleteBlog(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	blog := new(model.Blog)
	blog.Id = ParseInt(id)
	err := blog.Delete()
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func DeleteBlogArray(ctx *middleware.Context) {
	blogArray := ctx.R.FormValue("blogArray")
	blog := new(model.Blog)
	var res []int
	json.Unmarshal([]byte(blogArray), &res)
	err := blog.DeleteBlogArray(res)
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}
