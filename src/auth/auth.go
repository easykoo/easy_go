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
					if CheckPermission(user.(model.User), req.(int)) {
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

func CheckPermission(user model.User, module int) bool {
	privilege := &model.Privilege{ModuleId: module, RoleId: user.RoleId, DeptId: user.DeptId}
	exist, err := privilege.CheckModulePrivilege()
	PanicIf(err)
	return exist
}
