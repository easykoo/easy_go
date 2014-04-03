package handler

import (
	"github.com/martini-contrib/binding"
	"github.com/qiniu/log"

	"middleware"
	"model"
	"util"
)

func LogoutHandler(resp middleware.Response) {
	resp.Sn().Set("SignedUser", nil)
	resp.SessionSet("SignedUser", nil)
	resp.Render().HTML(200, "index", resp)
}

func LoginHandler(resp middleware.Response, formErr binding.Errors, loginUser model.UserLoginForm) {
	switch resp.Req().Method {
	case "POST":
		resp.JoinFormErrors(formErr)
		user := &model.User{Username: loginUser.Username, Password: loginUser.Password}
		if !resp.HasError() {
			if has, err := user.Exist(); has {
				util.PanicIf(err)
				var result *model.User
				result, err = user.GetUser()
				util.PanicIf(err)
				resp.Sn().Set("SignedUser", result)
				resp.SessionSet("SignedUser", result)
				resp.RegisterSessionProperty("SignedUser")
				var users []model.User
				users, err = user.SelectAll()
				util.PanicIf(err)
				resp.Set("users", users)
				log.Debug(result.Username, "login")
				resp.Render().Redirect("/admin/dashboard", 302)
			} else {
				resp.Set("user", user)
				resp.AddError("invalid username or password")
				resp.Render().HTML(200, "user/login", resp)
			}
		} else {
			resp.Render().HTML(200, "user/login", resp)
		}
	default:
		resp.Render().HTML(200, "user/login", resp)
	}
}

func RegisterHandler(resp middleware.Response, formErr binding.Errors, user model.UserRegisterForm) {
	switch resp.Req().Method {
	case "POST":
		log.Println(formErr)
		resp.SetFormErrors(formErr)
		log.Println(resp.FieldErrors())
		if !resp.HasError() {
			dbUser := model.User{Username: user.Username, Password: user.Password, Email: user.Email}

			if exist, err := dbUser.ExistUsername(); exist {
				util.PanicIf(err)
				resp.AddFieldError("username", "This username already exists.")
			}

			if exist, err := dbUser.ExistEmail(); exist {
				util.PanicIf(err)
				resp.AddFieldError("email", "This email already exists.")
			}

			log.Println(resp.FieldErrors())
			if !resp.HasError() {
				err := dbUser.Insert()
				util.PanicIf(err)
				resp.AddMessage("Register successfully!")
			} else {
				resp.Set("user", user)
			}
			resp.Render().HTML(200, "user/register", resp)
		}else {
			resp.Set("user", user)
			resp.Render().HTML(200, "user/register", resp)
		}
	default:
		resp.Render().HTML(200, "user/register", resp)
	}
}
