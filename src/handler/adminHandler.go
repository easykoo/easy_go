package handler

import (
	"middleware"
)

func DashboardHandler(resp middleware.Response) {
	resp.Render().HTML(200, "admin/dashboard", resp)
}
