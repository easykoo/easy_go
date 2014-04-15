package handler

import (
	"github.com/go-martini/martini"

	. "common"
	"middleware"
	"model"

	"encoding/json"
	"strconv"
	"time"
)

func PublishBlog(ctx *middleware.Context, blog model.Blog) {
	switch ctx.R.Method {
	case "POST":
		if blog.Title == "" || blog.Content == "" {
			ctx.AddError(Translate(ctx.SessionGet("Lang").(string), "message.error.publish.failed"))
		} else {
			blog.UpdateUser = ctx.SessionGet("SignedUser").(model.User).Username
			blog.State = "PUBLISHED"
			blog.PublishDate = time.Now()
			if blog.Version == 0 {
				blog.Priority = 5
				blog.CreateUser = ctx.SessionGet("SignedUser").(model.User).Username
				err := blog.Insert()
				PanicIf(err)
			} else {
				err := blog.Update()
				PanicIf(err)
			}
		}
		ctx.Redirect("/blog/view/" + strconv.Itoa(blog.Id))
	default:
		ctx.HTML(200, "blog/edit", ctx)
	}
}

func SaveBlog(ctx *middleware.Context, blog model.Blog) {
	if blog.Title == "" || blog.Content == "" {
		ctx.AddError(Translate(ctx.SessionGet("Lang").(string), "message.error.save.failed"))
	} else {
		blog.UpdateUser = ctx.SessionGet("SignedUser").(model.User).Username
		if blog.Version == 0 {
			blog.State = "DRAFT"
			blog.Priority = 5
			blog.CreateUser = ctx.SessionGet("SignedUser").(model.User).Username
			err := blog.Insert()
			PanicIf(err)
		} else {
			err := blog.Update()
			PanicIf(err)
		}
		dbBlog, err := blog.GetBlogById()
		PanicIf(err)
		ctx.Set("Blog", dbBlog)
		ctx.AddMessage(Translate(ctx.SessionGet("Lang").(string), "message.save.success"))
	}
	ctx.HTML(200, "blog/edit", ctx)
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
		blogList, total, err := blog.SearchByPage(false)
		PanicIf(err)
		ctx.Set("aaData", blogList)
		ctx.Set("iTotalDisplayRecords", total)
		ctx.Set("iTotalRecords", total)
		ctx.JSON(200, ctx.Response)
	default:
		ctx.HTML(200, "blog/allBlog", ctx)
	}
}

func Blog(ctx *middleware.Context) {
	blog := new(model.Blog)
	blog.SetPageActive(true)
	blog.SetPageSize(10)
	blog.SetDisplayStart(ParseInt(ctx.R.FormValue("iDisplayStart")))
	blog.AddSortProperty("publish_date", "desc")
	blogList, total, err := blog.SearchByPage(true)
	PanicIf(err)
	ctx.Set("BlogList", blogList)
	ctx.Set("Total", total)
	ctx.HTML(200, "blog", ctx)
}

func ViewBlog(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	blog := new(model.Blog)
	blog.Id = ParseInt(id)
	err := blog.GetBlog()
	PanicIf(err)
	ctx.Set("Blog", blog)
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

func EditBlog(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	blog := new(model.Blog)
	blog.Id = ParseInt(id)
	err := blog.GetBlog()
	PanicIf(err)
	ctx.Set("Blog", blog)
	ctx.HTML(200, "blog/edit", ctx)
}
