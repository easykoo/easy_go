package auth

import (
	"github.com/go-martini/martini"

	"middleware"
	."common"
	"reflect"
	"model"
)

const (
	SignInRequire  = "SIGNIN"
	SignOutRequire = "SIGNOUT"
	AdminRequire   = "ADMIN"
	Url            = "URL"
	Module_Admin   = 1
	Module_Account = 2
	Module_Profile = 3
)

func AuthRequest(privilege interface{}) martini.Handler {
	return func(ctx *middleware.Context) {
		Log.Info(ctx.R.RequestURI)
		switch privilege{
		case SignInRequire:
			if user := ctx.SessionGet("SignedUser"); user != nil {
				return
			}
			ctx.Redirect("/user/login")
			return
		case SignOutRequire:
			if user := ctx.SessionGet("SignedUser"); user == nil {
				return
			}
			ctx.Redirect("/user/login")
			return
		case AdminRequire:
			if user := ctx.SessionGet("SignedUser"); user != nil {
				if user.(model.User).RoleId == 1 {
					return
				}
			}
			ctx.HTML(403, "error/403", ctx)
			return
		default:
			if reflect.TypeOf(privilege).Kind() == reflect.Int {
				if privilege == 1 {
					return
				}
			}
			ctx.HTML(403, "error/403", ctx)
			return
		}
	}
}

/*

func Toggle(options *ToggleOptions) martini.Handler {
	return func(ctx *Context) {
		if options.SignOutRequire && ctx.IsSigned && ctx.Req.RequestURI != "/" {
			ctx.Redirect("/")
			return
		}

		if !options.DisableCsrf {
			if ctx.Req.Method == "POST" {
				if !ctx.CsrfTokenValid() {
					ctx.Error(403, "CSRF token does not match")
					return
				}
			}
		}

		if options.SignInRequire {
			if !ctx.IsSigned {
				ctx.Render().SetCookie("redirect_to", "/"+url.QueryEscape(ctx.Req().RequestURI))
				ctx.Redirect("/user/login")
				return
			} else if !ctx.User.IsActive{
				ctx.Data["Title"] = "Activate Your Account"
				ctx.HTML(200, "user/active")
				return
			}
		}

		if options.AdminRequire {
			if !ctx.User.IsAdmin {
				ctx.Error(403)
				return
			}
			ctx.Data["PageIsAdmin"] = true
		}
	}
}
*/
