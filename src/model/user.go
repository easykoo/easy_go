package model

import (
	"github.com/martini-contrib/binding"
	"github.com/qiniu/log"
	"net/http"
	"time"
	"regexp"
)

type User struct {
	Id         int       `form:"id" xorm:"int(11) not null autoincr"`
	Username   string    `form:"username" xorm:"varchar(20) not null"`
	Password   string    `form:"password" xorm:"varchar(60) not null"`
	FullName   string    `form:"fullName" xorm:"varchar(20) null"`
	Gender     int       `form:"gender" xorm:"int(1) "`
	Qq         int       `form:"qq" xorm:"int(16) default null"`
	Tel        string    `form:"tel" xorm:"varchar(20) null"`
	Postcode   string    `form:"postcode" xorm:"varchar(10) default null"`
	Address    string    `form:"address" xorm:"varchar(80) default null"`
	Email      string    `form:"email" json:"email" xorm:"varchar(45) default null"`
	RoleId     int       `xorm:"int(3) not null default 3"`
	DeptId     int       `xorm:"int(3) not null default 1"`
	Active     bool      `xorm:"tinyint(1) default 0"`
	Locked     bool      `xorm:"tinyint(1) default 0"`
	FailTime   int       `xorm:"int(1) default 0"`
	EffectDate time.Time `xorm:"datetime default null"`
	CreateUser string    `xorm:"varchar(20) default null"`
	CreateDate time.Time `xorm:"datetime created"`
	UpdateUser string    `xorm:"varchar(20) default null"`
	UpdateDate time.Time `xorm:"datetime updated"`
	Version    int       `xorm:"int(11) version"`
}

func (user *User) Exist() (bool, error) {
	return orm.Get(user)
}

func (user *User) ExistUsername() (bool, error) {
	return orm.Get(&User{Username:user.Username})
}

func (user *User) ExistEmail() (bool, error) {
	return orm.Get(&User{Email:user.Email})
}

func (user *User) GetUser() (*User, error) {
	_, err := orm.Get(user)
	return user, err
}

func (user *User) Save() error {
	_, err := orm.Insert(user)
	log.Println("user inserted")
	return err
}

func (user *User) SelectAll() ([]User, error) {
	var users []User
	err := orm.Find(&users)
	return users, err
}

type UserLoginForm struct {
	Username   string    `form:"username" binding:"required"`
	Password   string    `form:"password" binding:"required"`
}

func (user *UserLoginForm) Validate(errors *binding.Errors, req *http.Request) {
	if len(user.Username) < 5 {
		errors.Fields["username"] = "Length of username should be longer than 5."
	}
	if len(user.Password) < 5 {
		errors.Fields["password"] = "Length of password should be longer than 5."
	}
}

type UserRegisterForm struct {
	Username   string    `form:"username" binding:"required"`
	Password   string    `form:"password" binding:"required"`
	Email      string    `form:"email" binding:"required"`
}

func (user *UserRegisterForm) Validate(errors *binding.Errors, req *http.Request) {
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
