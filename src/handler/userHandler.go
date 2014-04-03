package handler

import (
	"github.com/martini-contrib/binding"
	"github.com/qiniu/log"

	"middleware"
	"model"
	"util"
)

func LogoutHandler(ctx *middleware.Context) {
	ctx.SessionSet("SignedUser", nil)
	ctx.HTML(200, "index", ctx)
}

func LoginHandler(ctx *middleware.Context, formErr binding.Errors, loginUser model.UserLoginForm) {
	switch ctx.Method {
	case "POST":
		ctx.JoinFormErrors(formErr)
		user := &model.User{Username: loginUser.Username, Password: loginUser.Password}
		if !ctx.HasError() {
			if has, err := user.Exist(); has {
				util.PanicIf(err)
				var result *model.User
				result, err = user.GetUser()
				util.PanicIf(err)
				ctx.SessionSet("SignedUser", result)
				var users []model.User
				users, err = user.SelectAll()
				util.PanicIf(err)
				ctx.Set("users", users)
				log.Debug(result.Username, "login")
				ctx.Redirect("/admin/dashboard")
			} else {
				ctx.Set("user", user)
				ctx.AddError("invalid username or password")
				ctx.HTML(200, "user/login", ctx)
			}
		} else {
			ctx.HTML(200, "user/login", ctx)
		}
	default:
		ctx.HTML(200, "user/login", ctx)
	}
}

func RegisterHandler(ctx *middleware.Context, formErr binding.Errors, user model.UserRegisterForm) {
	switch ctx.Method {
	case "POST":
		ctx.JoinFormErrors(formErr)
		if !ctx.HasError() {
			dbUser := model.User{Username: user.Username, Password: user.Password, Email: user.Email}

			if exist, err := dbUser.ExistUsername(); exist {
				util.PanicIf(err)
				ctx.AddFieldError("username", "This username already exists.")
			}

			if exist, err := dbUser.ExistEmail(); exist {
				util.PanicIf(err)
				ctx.AddFieldError("email", "This email already exists.")
			}

			if !ctx.HasError() {
				err := dbUser.Insert()
				util.PanicIf(err)
				ctx.AddMessage("Register successfully!")
			} else {
				ctx.Set("user", user)
			}
			ctx.HTML(200, "user/register", ctx)
		}else {
			ctx.Set("user", user)
			ctx.HTML(200, "user/register", ctx)
		}
	default:
		ctx.HTML(200, "user/register", ctx)
	}
}
