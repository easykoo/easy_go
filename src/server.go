package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	. "auth"
	. "common"
	"handler"
	"middleware"
	"model"

	"encoding/gob"
	"html/template"
	"os"
	"time"
)

func init() {
	gob.Register(model.User{})
	Log.Debug("server initializing...")
}

func newMartini() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(middleware.GetLogger())
	m.Map(model.SetEngine())
	m.Use(martini.Recovery())
	m.Use(martini.Static("public"))
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	m.Use(sessions.Sessions("my_session", middleware.NewFileStore(60*30)))
	//m.Use(sessions.Sessions("my_session", middleware.NewMemoryStore(60*30)))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		Funcs: []template.FuncMap{
			{
				"formatTime": func(args ...interface{}) string {
					t1 := time.Unix(args[0].(int64), 0)
					return t1.Format(time.Stamp)
				},
				"unescaped": func(args ...interface{}) template.HTML {
					return template.HTML(args[0].(string))
				},
				"equal": func(args ...interface{}) bool {
					return args[0] == args[1]
				},
				"tsl": func(lang string, format string) string {
					return Translate(lang, format)
				},
				"tslf": func(lang string, format string, args ...interface{}) string {
					return Translatef(lang, format, args...)
				},
				"privilege": func(user model.User, module int) bool {
					return CheckPermission(user, module)
				},
			},
		},
	}))

	m.Use(middleware.InitContext())

	return &martini.ClassicMartini{m, r}
}

func main() {
	m := newMartini()

	m.Get("/", handler.IndexHandler)
	m.Get("/index", handler.IndexHandler)
	m.Post("/contact", binding.Form(model.Feedback{}), handler.Contact)
	m.Get("/language/change/:lang", handler.LangHandler)

	m.Group("/user", func(r martini.Router) {
		r.Any("", AuthRequest(Module_Account), handler.AllUserHandler)
		r.Any("/logout", handler.LogoutHandler)
		r.Any("/login", binding.Form(model.UserLoginForm{}), handler.LoginHandler)
		r.Any("/register", binding.Form(model.UserRegisterForm{}), handler.RegisterHandler)
		r.Any("/delete", AuthRequest(Module_Account), handler.DeleteUsers)
		r.Any("/delete/:id", AuthRequest(Module_Account), handler.DeleteUser)
		r.Any("/role", AuthRequest(Module_Account), handler.SetRole)
		r.Any("/ban/:id", AuthRequest(Module_Account), handler.BanUser)
		r.Any("/lift/:id", AuthRequest(Module_Account), handler.LiftUser)
		//		r.Any("/:id", GetUser)
		//		r.Any("/new", NewUser)
		//		r.Any("/update/:id", UpdateUser)
	})

	m.Group("/profile", func(r martini.Router) {
		r.Any("/profile", AuthRequest(SignInRequired), binding.Form(model.User{}), handler.ProfileHandler)
		r.Any("/password", AuthRequest(SignInRequired), binding.Form(model.Password{}), handler.PasswordHandler)
		r.Any("/checkEmail", AuthRequest(SignInRequired), binding.Form(model.User{}), handler.CheckEmail)
	})

	m.Group("/admin", func(r martini.Router) {
		r.Get("/dashboard", AuthRequest(SignInRequired), handler.DashboardHandler)
		r.Get("/settings", AuthRequest(Module_Admin), handler.DashboardHandler)
	})

	m.Group("/feedback", func(r martini.Router) {
			r.Any("", AuthRequest(Module_Feedback), handler.AllFeedback)
			r.Any("/info", AuthRequest(Module_Feedback), handler.FeedbackInfo)
			r.Any("/delete", AuthRequest(Module_Feedback), handler.DeleteFeedbackArray)
			r.Any("/delete/:id", AuthRequest(Module_Feedback), handler.DeleteFeedback)
			r.Any("/view/:id", AuthRequest(Module_Feedback), handler.ViewFeedback)
		})

	Log.Info("server is started...")
	os.Setenv("PORT", Cfg.MustValue("", "http_port"))
	m.Run()
}
