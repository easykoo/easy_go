package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/qiniu/log"

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
	log.SetOutputLevel(log.Ldebug)
	log.Debug("server initializing...")
}

func newMartini() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	m.Use(martini.Logger())
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

	m.Use(middleware.InitResponse())

	return &martini.ClassicMartini{m, r}
}

func main() {
	m := newMartini()

	m.Get("/", handler.IndexHandler)
	m.Get("/index", handler.IndexHandler)

	m.Group("/user", func(r martini.Router) {
		r.Any("/logout", handler.LogoutHandler)
		r.Any("/login", binding.Form(model.UserLoginForm{}), handler.LoginHandler)
		r.Any("/register", binding.Form(model.UserRegisterForm{}), handler.RegisterHandler)
		//		r.Get("/:id", GetBooks)
		//		r.Post("/new", NewBook)
		//		r.Put("/update/:id", UpdateBook)
		//		r.Delete("/delete/:id", DeleteBook)
	})

	m.Group("/admin", func(r martini.Router) {
		r.Get("/dashboard", handler.DashboardHandler)
	})

	log.Info("server is started...")
	os.Setenv("PORT", "80")
	m.Run()
}
