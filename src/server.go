package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"handler"
	"middleware"
	"model"
	"auth"
	. "common"

	"encoding/gob"
	"html/template"
	"os"
	"time"
)

func init() {
	gob.Register(model.User{})
	Log.Debug("server initializing...")
	Log.Debugf("server initializing...%d",50)
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

	m.Group("/user", func(r martini.Router) {
		r.Any("/logout", auth.AuthRequest(auth.SignInRequire), handler.LogoutHandler)
		r.Any("/login", auth.AuthRequest(auth.SignOutRequire), binding.Form(model.UserLoginForm{}), handler.LoginHandler)
		r.Any("/register", auth.AuthRequest(auth.SignOutRequire), binding.Form(model.UserRegisterForm{}), handler.RegisterHandler)
		//		r.Get("/:id", GetBooks)
		//		r.Post("/new", NewBook)
		//		r.Put("/update/:id", UpdateBook)
		//		r.Delete("/delete/:id", DeleteBook)
	})

	m.Group("/admin", func(r martini.Router) {
			r.Get("/dashboard", auth.AuthRequest(auth.SignInRequire), handler.DashboardHandler)
			r.Get("/settings", auth.AuthRequest(auth.Module_Account), handler.DashboardHandler)
	})

	Log.Info("server is started...")
	os.Setenv("PORT", "80")
	m.Run()
}
