package middleware

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"net/http"
)

var sessionProperties []string

type Response interface {
	Get(key interface{}) interface{}
	Set(key interface{}, val interface{})
	SessionSet(key interface{}, val interface{})
	Delete(key interface{})
	Clear()
	AddMessage(message string)
	ClearMessages()
	HasMessage() bool
	SetErrors(err binding.Errors)
	JoinErrors(err binding.Errors)
	AddError(err string)
	AddFieldError(field string, err string)
	ClearError()
	HasError() bool
	HasCommonError() bool
	HasFieldError() bool
	HasOverallError() bool
	OverallErrors() map[string]string
	FieldErrors() map[string]string
	initIf()
	Render() render.Render
	Req() *http.Request
	Sn() sessions.Session
	RegisterSessionProperty(property string)
	TransferSessionProperties(s sessions.Session)
}

type response struct {
	Response     map[interface{}]interface{}
	Session      map[interface{}]interface{}
	Rnd          render.Render
	R            *http.Request
	S            sessions.Session
	FormErr      binding.Errors
	Messages     []string
	CommonErrors []string
}

func NewResponse() *response {
	return &response{
		Session:  make(map[interface{}]interface{}),
		Response: make(map[interface{}]interface{}),
		FormErr:  *&binding.Errors{Overall: make(map[string]string), Fields: make(map[string]string)},
	}
}

func (self *response) initIf() {
	if self.Response == nil {
		self.Response = make(map[interface{}]interface{})
	}
	if self.FormErr.Fields == nil {
		self.FormErr.Fields = make(map[string]string)
	}
	if self.FormErr.Overall == nil {
		self.FormErr.Overall = make(map[string]string)
	}
}

func (self *response) Get(key interface{}) interface{} {
	return self.Response[key]
}

func (self *response) Set(key interface{}, val interface{}) {
	self.initIf()
	self.Response[key] = val
}

func (self *response) SessionSet(key interface{}, val interface{}) {
	self.Session[key] = val
}

func (self *response) Delete(key interface{}) {
	delete(self.Response, key)
}

func (self *response) Clear() {
	for key := range self.Response {
		self.Delete(key)
	}
}

func (self *response) AddMessage(message string) {
	self.Messages = append(self.Messages, message)
}

func (self *response) ClearMessages() {
	self.Messages = self.Messages[:0]
}

func (self *response) HasMessage() bool {
	return (len(self.Messages) > 0)
}

func (self *response) SetErrors(err binding.Errors) {
	self.FormErr = err
}

func (self *response) JoinErrors(err binding.Errors) {
	self.initIf()
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

func (self *response) AddError(err string) {
	self.CommonErrors = append(self.CommonErrors, err)
}

func (self *response) AddFieldError(field string, err string) {
	self.FormErr.Fields[field] = err
}

func (self *response) ClearError() {
	self.CommonErrors = self.CommonErrors[:0]

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

func (self *response) HasError() bool {
	return (len(self.CommonErrors) > 0) || (len(self.FormErr.Overall) > 0) || (len(self.FormErr.Fields) > 0)
}

func (self *response) HasCommonError() bool {
	return (len(self.CommonErrors) > 0)
}

func (self *response) HasFieldError() bool {
	return (len(self.FormErr.Fields) > 0)
}

func (self *response) HasOverallError() bool {
	return (len(self.FormErr.Overall) > 0)
}

func (self *response) OverallErrors() map[string]string {
	return self.FormErr.Overall
}

func (self *response) FieldErrors() map[string]string {
	return self.FormErr.Fields
}

func (self *response) Render() render.Render {
	return self.Rnd
}

func (self *response) Req() *http.Request {
	return self.R
}

func (self *response) Sn() sessions.Session {
	return self.S
}

func (self *response) RegisterSessionProperty(property string) {
	for _, val := range sessionProperties {
		if val == property {
			return
		}
	}
	sessionProperties = append(sessionProperties, property)
}

func (self *response) TransferSessionProperties(s sessions.Session) {
	for _, val := range sessionProperties {
		self.SessionSet(val, s.Get(val))
	}
}

func InitResponse() martini.Handler {
	return func(c martini.Context, s sessions.Session, rnd render.Render, r *http.Request) {
		resp := NewResponse()
		resp.Rnd = rnd
		resp.R = r
		resp.S = s
		resp.TransferSessionProperties(s)
		c.MapTo(resp, (*Response)(nil))
	}
}
