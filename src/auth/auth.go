package auth

import (
	"github.com/go-martini/martini"

	. "common"
	"middleware"
	"model"
	"reflect"
)

const (
	SignInRequired = 9
	Module_Admin   = 1 + iota
	Module_Account
	Module_Feedback
	Module_News
	Module_Product
)

func AuthRequest(req interface{}) martini.Handler {
	return func(ctx *middleware.Context) {
		Log.Info("Checking privilege: ", ctx.R.RequestURI)
		switch req {
		case SignInRequired:
			Log.Info("Checking style: ", "SignInRequired")
			if user := ctx.SessionGet("SignedUser"); user != nil {
				return
			}
			ctx.Redirect("/user/login")
			return
		default:
			Log.Info("Checking style: ", "Module")
			if user := ctx.SessionGet("SignedUser"); user != nil {
				if reflect.TypeOf(req).Kind() == reflect.Int {
					privilege := &model.Privilege{ModuleId: req.(int), RoleId: user.(model.User).RoleId, DeptId: user.(model.User).DeptId}
					exist, err := privilege.CheckModulePrivilege()
					PanicIf(err)
					if exist {
						return
					}
					ctx.HTML(403, "error/403", ctx)
					return
				}
			} else {
				ctx.Redirect("/user/login")
				return
			}
			ctx.HTML(403, "error/403", ctx)
			return
		}
	}
}
