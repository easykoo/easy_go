package model

import (
	"time"
)

type Category struct {
	Id          int    `form:"id" xorm:"int(3) pk not null autoincr"`
	Description string `form:"description" xorm:"varchar(20) not null"`
	//	Parent      *Category `form:"parentId" xorm:"parent_id int(3)"`
	CreateUser string    `xorm:"varchar(20) default 'SYSTEM'"`
	CreateDate time.Time `xorm:"datetime created"`
	UpdateUser string    `xorm:"varchar(20) default 'SYSTEM'"`
	UpdateDate time.Time `xorm:"datetime updated"`
	Version    int       `form:"version" xorm:"int(11) version"`
}

func (self *Category) GetRoleById(id int) (*Category, error) {
	category := &Category{Id: id}
	_, err := orm.Get(category)
	return category, err
}
