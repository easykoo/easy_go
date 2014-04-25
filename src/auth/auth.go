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
	Module_Admin   = iota
	Module_Account
	Module_Feedback
	Module_News
	Module_Product
	Module_Blog
)

func AuthRequest(req interface{}) martini.Handler {
	return func(ctx *middleware.Context) {
		Log.Info("Checking privilege: ", ctx.R.RequestURI)
		switch req {
		case SignInRequired:
			Log.Info("Checking style: ", "SignInRequired")
			if user := ctx.SessionGet("SignedUser"); user != nil {
				Log.Info("Pass!")
				return
			}
			ctx.Redirect("/user/login")
			return
		default:
			Log.Info("Checking style: ", "Module ", req.(int))
			if user := ctx.SessionGet("SignedUser"); user != nil {
				if reflect.TypeOf(req).Kind() == reflect.Int {
					if CheckPermission(user, req.(int)) {
						Log.Info("Pass!")
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

func CheckPermission(user interface{}, module int) bool {
	Log.Debug("Type: ", reflect.TypeOf(user).Kind())
	if reflect.TypeOf(user).Kind() == reflect.Struct {
		val := user.(model.User)
		privilege := &model.Privilege{ModuleId: module, RoleId: val.Role.Id, DeptId: val.Dept.Id}
		exist, err := privilege.CheckModulePrivilege()
		PanicIf(err)
		return exist
	} else {
		val := user.(*model.User)
		privilege := &model.Privilege{ModuleId: module, RoleId: val.Role.Id, DeptId: val.Dept.Id}
		exist, err := privilege.CheckModulePrivilege()
		PanicIf(err)
		return exist
	}
}
