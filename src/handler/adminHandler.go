package handler

import "middleware"

func DashboardHandler(ctx *middleware.Context) {
	ctx.HTML(200, "admin/dashboard", ctx)
}

func SettingsHandler(ctx *middleware.Context) {
	ctx.HTML(200, "admin/settings", ctx)
}
