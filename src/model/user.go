package model

import (
	"github.com/martini-contrib/binding"

	. "common"

	"net/http"
	"regexp"
	"time"
)

type User struct {
	Id         int       `form:"id" xorm:"int(11) pk not null autoincr"`
	Username   string    `form:"username" xorm:"varchar(20) not null"`
	Password   string    `form:"password" xorm:"varchar(60) not null"`
	FullName   string    `form:"fullName" xorm:"varchar(20) null"`
	Gender     int       `form:"gender" xorm:"int(1) default 0"`
	Qq         string    `form:"qq" xorm:"varchar(16) default null"`
	Tel        string    `form:"tel" xorm:"varchar(20) null"`
	Postcode   string    `form:"postcode" xorm:"varchar(10) default null"`
	Address    string    `form:"address" xorm:"varchar(80) default null"`
	Email      string    `form:"email" xorm:"varchar(45) unique"`
	RoleId     int       `xorm:"int(3) default 3"`
	DeptId     int       `xorm:"int(3) default 1"`
	Active     bool      `xorm:"tinyint(1) default 0"`
	Locked     bool      `xorm:"tinyint(1) default 0"`
	FailTime   int       `xorm:"int(1) default 0"`
	CreateUser string    `json:"create_user" xorm:"varchar(20) default 'SYSTEM'"`
	CreateDate time.Time `json:"create_date" xorm:"datetime created"`
	UpdateUser string    `json:"update_user" xorm:"varchar(20) default 'SYSTEM'"`
	UpdateDate time.Time `json:"update_date" xorm:"datetime updated"`
	Version    int       `form:"version" xorm:"int(11) version"`
}

func (self *User) Exist() (bool, error) {
	return orm.Get(self)
}

func (self *User) ExistUsername() (bool, error) {
	return orm.Get(&User{Username: self.Username})
}

func (self *User) ExistEmail() (bool, error) {
	return orm.Get(&User{Email: self.Email})
}

func (self *User) GetUser() (*User, error) {
	_, err := orm.Id(self.Id).Get(self)
	return self, err
}

func (self *User) GetUserById(id int) (*User, error) {
	user := &User{Id: id}
	_, err := orm.Get(user)
	return user, err
}

func (self *User) Insert() error {
	self.DeptId = 1
	self.RoleId = 3
	self.CreateUser = "SYSTEM"
	self.UpdateUser = "SYSTEM"
	_, err := orm.InsertOne(self)
	Log.Info(self.Username, "inserted")
	return err
}

func (self *User) Update() error {
	_, err := orm.Id(self.Id).Update(self)
	Log.Info(self.Username, "updated")
	return err
}

func (self *User) Delete() error {
	_, err := orm.Delete(self)
	Log.Info(self.Username, "deleted")
	return err
}

func (self *User) SelectAll() ([]User, error) {
	var users []User
	err := orm.Find(&users)
	return users, err
}

type UserLoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserRegisterForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

func (user UserRegisterForm) Validate(errors *binding.Errors, r *http.Request) {
	if len(user.Username) < 5 {
		errors.Fields["username"] = "Length of username should be longer than 5."
	}
	if len(user.Password) < 5 {
		errors.Fields["password"] = "Length of password should be longer than 5."
	}
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(user.Email))
	if matched == false {
		errors.Fields["email"] = "Please enter a valid email address."
	}
}
