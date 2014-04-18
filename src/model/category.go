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
	Page       `xorm:"-"`
}

func (self *Category) GetRoleById(id int) (*Category, error) {
	category := &Category{Id: id}
	_, err := orm.Get(category)
	return category, err
}

func (self *Category) SearchByPage() ([]Category, int, error) {
	total, err := orm.Count(self)
	var category []Category

	session := orm.NewSession()
	defer session.Close()
	if len(self.GetSortProperties()) > 0 {
		session = session.OrderBy(self.GetSortProperties()[0].Column + " " + self.GetSortProperties()[0].Direction)
	}
	err = session.Limit(self.GetPageSize(), self.GetDisplayStart()).Find(&category, self)
	return category, int(total), err
}
