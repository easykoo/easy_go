package middleware

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	. "common"

	"net/http"
)

var sessionProperties []string

type Context struct {
	render.Render
	C        martini.Context
	S        sessions.Session
	R        *http.Request
	W        http.ResponseWriter
	FormErr  binding.Errors
	Messages []string
	Errors   []string
	Response map[string]interface{}
	Session  map[string]interface{}
}

func (self *Context) init() {
	if self.Response == nil {
		self.Response = make(map[string]interface{})
	}
	if self.Session == nil {
		self.Session = make(map[string]interface{})
	}
	if self.FormErr.Fields == nil {
		self.FormErr.Fields = make(map[string]string)
	}
	if self.FormErr.Overall == nil {
		self.FormErr.Overall = make(map[string]string)
	}
}

func (self *Context) SessionGet(key string) interface{} {
	return self.S.Get(key)
}

func (self *Context) SessionSet(key string, val interface{}) {
	self.init()
	self.S.Set(key, val)
	self.Session[key] = val
	for _, val := range sessionProperties {
		if val == key {
			return
		}
	}
	sessionProperties = append(sessionProperties, key)
}

func (self *Context) SessionDelete(key string) {
	delete(self.Response, key)
	self.S.Delete(key)
}

func (self *Context) SessionClear() {
	self.Clear()
	self.S.Clear()
}

func (self *Context) Get(key string) interface{} {
	return self.Response[key]
}

func (self *Context) Set(key string, val interface{}) {
	self.init()
	self.Response[key] = val
}

func (self *Context) Delete(key string) {
	delete(self.Response, key)
}

func (self *Context) Clear() {
	for key := range self.Response {
		self.Delete(key)
	}
}

func (self *Context) AddMessage(message string) {
	self.Messages = append(self.Messages, message)
}

func (self *Context) ClearMessages() {
	self.Messages = self.Messages[:0]
}

func (self *Context) HasMessage() bool {
	return (len(self.Messages) > 0)
}

func (self *Context) SetFormErrors(err binding.Errors) {
	self.FormErr = err
}

func (self *Context) JoinFormErrors(err binding.Errors) {
	self.init()
	for key, val := range err.Fields {
		if _, exists := self.FormErr.Fields[key]; !exists {
			self.FormErr.Fields[key] = val
		}
	}
	for key, val := range err.Overall {
		if _, exists := self.FormErr.Overall[key]; !exists {
			self.FormErr.Overall[key] = val
		}
	}
}

func (self *Context) AddError(err string) {
	self.Errors = append(self.Errors, err)
}

func (self *Context) AddFieldError(field string, err string) {
	self.FormErr.Fields[field] = err
}

func (self *Context) ClearError() {
	self.Errors = self.Errors[:0]

	for key, _ := range self.FormErr.Fields {
		if _, exists := self.FormErr.Fields[key]; exists {
			delete(self.FormErr.Fields, key)
		}
	}

	for key, _ := range self.FormErr.Overall {
		if _, exists := self.FormErr.Overall[key]; exists {
			delete(self.FormErr.Overall, key)
		}
	}
}

func (self *Context) HasError() bool {
	return self.HasCommonError() || self.HasFieldError() || self.HasOverallError()
}

func (self *Context) HasCommonError() bool {
	return (len(self.Errors) > 0)
}

func (self *Context) HasFieldError() bool {
	return (len(self.FormErr.Fields) > 0)
}

func (self *Context) HasOverallError() bool {
	return (len(self.FormErr.Overall) > 0)
}

func (self *Context) OverallErrors() map[string]string {
	return self.FormErr.Overall
}

func (self *Context) FieldErrors() map[string]string {
	return self.FormErr.Fields
}

func (self *Context) TransferSessionProperties(s sessions.Session) {
	lang := s.Get("Lang")
	if lang == nil {
		self.SessionSet("Lang", Cfg.MustValue("", "locale"))
	} else {
		self.SessionSet("Lang", lang)
	}
	for _, val := range sessionProperties {
		self.SessionSet(val, s.Get(val))
	}
}

func InitContext() martini.Handler {
	return func(c martini.Context, s sessions.Session, rnd render.Render, r *http.Request, w http.ResponseWriter) {
		ctx := &Context{
			Render: rnd,
			W:      w,
			R:      r,
			C:      c,
			S:      s,
		}
		ctx.TransferSessionProperties(s)
		c.Map(ctx)
	}
}
