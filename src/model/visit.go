package model

import (
	. "common"

	"strings"
	"time"
)

type Visit struct {
	Id         int       `form:"blogId" xorm:"int(11) pk not null autoincr"`
	SessionId  string    `xorm:"varchar(60) null"`
	Ip         string    `xorm:"varchar(15) null"`
	CreateUser string    `xorm:"varchar(20) default 'SYSTEM'"`
	CreateDate time.Time `xorm:"datetime created"`
	UpdateUser string    `xorm:"varchar(20) default 'SYSTEM'"`
	UpdateDate time.Time `xorm:"datetime updated"`
	Version    int       `form:"version" xorm:"int(11) version"`
	Page       `xorm:"-"`
}

func (self *Visit) SetIp(ip string) {
	ip = strings.Split(ip, ":")[0]
	if len(ip) >= 7 {
		self.Ip = ip
	}
}

func (self *Visit) Insert() error {
	_, err := orm.InsertOne(self)
	Log.Info("Visit ", self.Id, " ", self.SessionId, " inserted")
	return err
}

func (self *Visit) Update() error {
	_, err := orm.Update(self)
	Log.Info("Visit ", self.Id, " ", self.SessionId, " updated")
	return err
}

func (self *Visit) Delete() error {
	_, err := orm.Delete(self)
	Log.Info("Visit ", self.Id, " ", self.SessionId, " deleted")
	return err
}

func (self *Visit) SearchByPage() ([]Visit, int64, error) {
	total, err := orm.Count(self)
	var visit []Visit
	err = orm.OrderBy(self.GetSortProperties()[0].Column+" "+self.GetSortProperties()[0].Direction).Limit(self.GetPageSize(), self.GetDisplayStart()).Find(&visit, self)
	return visit, total, err
}
