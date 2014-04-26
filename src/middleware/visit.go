package middleware

import (
	"github.com/go-martini/martini"

	"github.com/martini-contrib/sessions"
	"model"
	"net/http"
	"reflect"
)

func RecordVisit() martini.Handler {
	return func(s sessions.Session, r *http.Request) {
		visit := new(model.Visit)
		visit.SessionId = s.GetId()
		user := s.Get("SignedUser")
		var id int
		if user != nil {
			if reflect.TypeOf(user).Kind() == reflect.Struct {
				id = user.(model.User).Id
			} else {
				id = user.(*model.User).Id
			}
		}
		visit.User = model.User{Id: id}
		visit.SetIp(r.RemoteAddr)
		if visit.ExistVisit() {
			visit.Update()
		} else {
			visit.Insert()
		}
	}
}
