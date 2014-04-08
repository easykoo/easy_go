package handler

import (
	"github.com/martini-contrib/binding"

	. "common"
	"middleware"
	"model"
)

func LogoutHandler(ctx *middleware.Context) {
	ctx.SessionSet("SignedUser", nil)
	ctx.HTML(200, "index", ctx)
}

func LoginHandler(ctx *middleware.Context, formErr binding.Errors, loginUser model.UserLoginForm) {
	switch ctx.R.Method {
	case "POST":
		ctx.JoinFormErrors(formErr)
		user := &model.User{Username: loginUser.Username, Password: loginUser.Password}
		if !ctx.HasError() {
			if has, err := user.Exist(); has {
				PanicIf(err)
				PanicIf(err)
				ctx.SessionSet("SignedUser", user)
				var users []model.User
				users, err = user.SelectAll()
				PanicIf(err)
				ctx.Set("users", users)
				Log.Info(user.Username, " login")
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
	switch ctx.R.Method {
	case "POST":
		ctx.JoinFormErrors(formErr)
		if !ctx.HasError() {
			dbUser := model.User{Username: user.Username, Password: user.Password, Email: user.Email}

			if exist, err := dbUser.ExistUsername(); exist {
				PanicIf(err)
				ctx.AddFieldError("username", "This username already exists.")
			}

			if exist, err := dbUser.ExistEmail(); exist {
				PanicIf(err)
				ctx.AddFieldError("email", "This email already exists.")
			}

			if !ctx.HasError() {
				err := dbUser.Insert()
				PanicIf(err)
				ctx.AddMessage("Register successfully!")
			} else {
				ctx.Set("user", user)
			}
			ctx.HTML(200, "user/register", ctx)
		} else {
			ctx.Set("user", user)
			ctx.HTML(200, "user/register", ctx)
		}
	default:
		ctx.HTML(200, "user/register", ctx)
	}
}

func ProfileHandler(ctx *middleware.Context, formErr binding.Errors, user model.User) {
	switch ctx.R.Method {
	case "POST":
		ctx.JoinFormErrors(formErr)
		if !ctx.HasError() {
			err := user.Update()
			PanicIf(err)
			dbUser, err := user.GetUserById(user.Id)
			PanicIf(err)
			ctx.AddMessage("Profile changed successfully!")
			ctx.SessionSet("SignedUser", dbUser)
		}
		ctx.HTML(200, "profile/profile", ctx)
	default:
		ctx.HTML(200, "profile/profile", ctx)
	}
}

func AllUserHandler(ctx *middleware.Context, user model.User) {
	switch ctx.R.Method {
	case "POST":
		users, err := user.SelectAll()
		PanicIf(err)
		ctx.Set("aaData", users);
		ctx.Set("iTotalDisplayRecords", len(users));
		ctx.Set("iTotalRecords", len(users));
		ctx.JSON(200, ctx.Response)
	default:
		ctx.HTML(200, "user/allUser", ctx)
	}
}
