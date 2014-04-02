package handler


import (
	"middleware"
)

func IndexHandler(resp middleware.Response) {
	resp.Render().HTML(200, "index", resp)
}

func HeadHandler(resp middleware.Response) {
	resp.Render().HTML(200, "layout/head", resp)
}

func LeftHandler(resp middleware.Response) {
	resp.Render().HTML(200, "layout/left", resp)
}
