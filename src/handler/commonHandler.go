package handler


import (
	"middleware"
)

func IndexHandler(resp middleware.Response) {
	resp.Render().HTML(200, "index", resp)
}
