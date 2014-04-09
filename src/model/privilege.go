package model

import "time"

type Privilege struct {
	ModuleId   int       `xorm:"int(11) default 3"`
	RoleId     int       `xorm:"int(11) default 3"`
	DeptId     int       `xorm:"int(11) default 3"`
	CreateUser string    `json:"create_user" xorm:"varchar(20) default 'SYSTEM'"`
	CreateDate time.Time `json:"create_date" xorm:"datetime created"`
	UpdateUser string    `json:"update_user" xorm:"varchar(20) default 'SYSTEM'"`
	UpdateDate time.Time `json:"update_date" xorm:"datetime updated"`
}

type Module struct {
	Id          int       `xorm:"int(3) pk not null"`
	description string    `xorm:"varchar(40) not null"`
	CreateUser  string    `json:"create_user" xorm:"varchar(20) default 'SYSTEM'"`
	CreateDate  time.Time `json:"create_date" xorm:"datetime created"`
	UpdateUser  string    `json:"update_user" xorm:"varchar(20) default 'SYSTEM'"`
	UpdateDate  time.Time `json:"update_date" xorm:"datetime updated"`
}

func (self *Privilege) CheckModulePrivilege() (bool, error) {
	privilege := &Privilege{ModuleId: self.ModuleId, RoleId: self.RoleId, DeptId: self.DeptId}
	return orm.Get(privilege)
}
