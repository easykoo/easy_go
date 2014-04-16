package model

import (
	. "common"

	"strconv"
	"time"
)

type Blog struct {
	Id          int       `form:"id" xorm:"int(11) pk not null autoincr"`
	CategoryId  int       `form:"category_id" xorm:"int(3) default null"`
	Title       string    `form:"title" xorm:"varchar(45) not null"`
	Content     string    `form:"content" xorm:"blob not null"`
	State       string    `xorm:"varchar(10) default null"`
	Priority    int       `xorm:"int(1) default 5"`
	PublishDate time.Time `json:"publish_date" xorm:"datetime default null"`
	CreateUser  string    `json:"create_user" xorm:"varchar(20) default null"`
	CreateDate  time.Time `json:"create_date" xorm:"datetime created"`
	UpdateUser  string    `json:"update_user" xorm:"varchar(20) default null"`
	UpdateDate  time.Time `json:"update_date" xorm:"datetime updated"`
	Version     int       `form:"version" xorm:"int(11) version"`
	Page        `xorm:"-"`
}

func (self *Blog) Insert() error {
	_, err := orm.InsertOne(self)
	Log.Info("Blog ", self.Id, " inserted")
	return err
}

func (self *Blog) Update() error {
	_, err := orm.Id(self.Id).Update(self)
	Log.Info("Blog ", self.Id, " updated!")
	return err
}

func (self *Blog) Delete() error {
	_, err := orm.Delete(self)
	Log.Info("Blog ", self.Id, " deleted")
	return err
}

func (self *Blog) GetBlogById() (*Blog, error) {
	blog := &Blog{Id: self.Id}
	_, err := orm.Get(blog)
	return blog, err
}

func (self *Blog) GetBlog() error {
	_, err := orm.Id(self.Id).Get(self)
	return err
}

func (self *Blog) SetState(state string) error {
	var err error
	_, err = orm.Id(self.Id).MustCols("state").Update(&Blog{State: state})
	return err
}

func (self *Blog) DeleteBlogArray(array []int) error {
	_, err := orm.In("id", array).Delete(&Blog{})
	sql := "delete from `blog` where id in ("
	for index, val := range array {
		sql += strconv.Itoa(val)
		if index < len(array)-1 {
			sql += ","
		}
	}
	sql += ")"
	_, err = orm.Exec(sql)
	Log.Info("Blog array: ", array, " deleted")
	return err
}

func (self *Blog) SearchByPage(content bool) ([]Blog, int, error) {
	total, err := orm.Count(self)
	var blog []Blog
	if content {
		err = orm.OrderBy(self.GetSortProperties()[0].Column+" "+self.GetSortProperties()[0].Direction).Limit(self.GetPageSize(), self.GetDisplayStart()).Find(&blog, self)
	} else {
		err = orm.Omit("content").OrderBy(self.GetSortProperties()[0].Column+" "+self.GetSortProperties()[0].Direction).Limit(self.GetPageSize(), self.GetDisplayStart()).Find(&blog, self)
	}
	return blog, int(total), err
}
