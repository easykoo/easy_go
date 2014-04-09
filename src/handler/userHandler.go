package handler

import (
	"github.com/martini-contrib/binding"

	. "common"
	"github.com/go-martini/martini"
	"middleware"
	"model"

	"encoding/json"
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

func AllUserHandler(ctx *middleware.Context) {
	switch ctx.R.Method {
	case "POST":
		user := new(model.User)
		user.SetPageActive(true)
		user.SetPageSize(ParseInt(ctx.R.FormValue("iDisplayLength")))
		user.SetDisplayStart(ParseInt(ctx.R.FormValue("iDisplayStart")))
		columnNum := ctx.R.FormValue("iSortCol_0")
		sortColumn := ctx.R.FormValue("mDataProp_" + columnNum)
		user.AddSortProperty(sortColumn, ctx.R.FormValue("sSortDir_0"))
		users, total, err := user.SearchByPage()
		PanicIf(err)
		ctx.Set("aaData", users)
		ctx.Set("iTotalDisplayRecords", total)
		ctx.Set("iTotalRecords", total)
		ctx.JSON(200, ctx.Response)
	default:
		ctx.HTML(200, "user/allUser", ctx)
	}
}

func DeleteUser(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	user := new(model.User)
	user.Id = ParseInt(id)
	err := user.Delete()
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func DeleteUsers(ctx *middleware.Context) {
	users:=ctx.R.FormValue("users")
	var res []int
	json.Unmarshal([]byte(users), &res)
	user := new(model.User)
	err := user.DeleteUsers(res)
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func AdminUser(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	user := new(model.User)
	user.Id = ParseInt(id)
	err := user.SetRole(1)
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func HireUser(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	user := new(model.User)
	user.Id = ParseInt(id)
	err := user.SetRole(2)
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func FireUser(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	user := new(model.User)
	user.Id = ParseInt(id)
	err := user.SetRole(3)
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func BanUser(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	user := new(model.User)
	user.Id = ParseInt(id)
	err := user.SetLock(true)
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}

func LiftUser(ctx *middleware.Context, params martini.Params) {
	id := params["id"]
	user := new(model.User)
	user.Id = ParseInt(id)
	err := user.SetLock(false)
	PanicIf(err)
	ctx.Set("success", true)
	ctx.JSON(200, ctx.Response)
}
