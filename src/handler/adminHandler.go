package handler

import (
	. "common"
	"middleware"
	"model"
)

func DashboardHandler(ctx *middleware.Context) {
	visit := new(model.Visit)
	visit.SetPageActive(true)
	visit.SetPageSize(10)
	pageNo := ParseInt(ctx.R.FormValue("page"))
	visit.SetPageNo(pageNo)
	visit.AddSortProperty("create_date", "desc")
	visitList, total, err := visit.SearchByPage()
	PanicIf(err)

	visit.SetTotalRecord(total)
	visit.Result = visitList
	ctx.Set("Visit", visit)
	ctx.HTML(200, "admin/dashboard", ctx)
}

func SettingsHandler(ctx *middleware.Context, settings model.Settings) {
	if ctx.R.Method == "POST" {
		err := settings.Update()
		PanicIf(err)
		dbSettings := model.GetSettings()
		ctx.AddMessage(Translate(ctx.SessionGet("Lang").(string), "message.change.success"))
		ctx.SessionSet("Settings", dbSettings)
	}
	user := &model.User{}
	users, err := user.SelectAll()
	PanicIf(err)
	ctx.Set("Users", users)

	ctx.HTML(200, "admin/settings", ctx)
}

func AboutHandler(ctx *middleware.Context) {
	settings := model.GetSettings()
	about := ctx.R.FormValue("about")
	settings.About = about
	err := settings.Update()
	PanicIf(err)
	dbSettings := model.GetSettings()
	ctx.SessionSet("Settings", dbSettings)

	ctx.Redirect("/about")
}
