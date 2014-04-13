package model

import (
	. "common"

	"strconv"
	"time"
)

type Feedback struct {
	Id         int       `xorm:"int(11) pk not null autoincr"`
	Name       string    `form:"name" xorm:"varchar(20) not null"`
	Email      string    `form:"email" xorm:"varchar(45) unique"`
	Content    string    `form:"content" xorm:"varchar(45) unique"`
	Viewed     bool      `xorm:"tinyint(1) default 0"`
	CreateDate time.Time `json:"create_date" xorm:"datetime created"`
	ViewDate   time.Time `json:"view_date" xorm:"datetime updated"`
	Page       `xorm:"-"`
}

func (self *Feedback) Insert() error {
	_, err := orm.InsertOne(self)
	Log.Info("Feedback ", self.Id, " inserted")
	return err
}

func (self *Feedback) Delete() error {
	_, err := orm.Delete(self)
	Log.Info("Feedback ", self.Id, " deleted")
	return err
}

func (self *Feedback) SetViewed(view bool) error {
	var err error
	_, err = orm.Id(self.Id).UseBool("viewed").Update(&Feedback{Viewed: view})
	return err
}

func (self *Feedback) DeleteFeedbackArray(array []int) error {
	_, err := orm.In("id", array).Delete(&Feedback{})
	sql := "delete from `feedback` where id in ("
	for index, val := range array {
		sql += strconv.Itoa(val)
		if index < len(array)-1 {
			sql += ","
		}
	}
	sql += ")"
	_, err = orm.Exec(sql)
	Log.Info("Feedback array: ", array, " deleted")
	return err
}

func (self *Feedback) Info() ([]Feedback, int, error) {
	var totalRecords int
	var feedback []Feedback
	err := orm.UseBool("viewed").MustCols("viewed").Find(&feedback, self)
	totalRecords = len(feedback)
	feedback = feedback[:0]
	err = orm.UseBool("viewed").MustCols("viewed").OrderBy("create_date desc").Limit(5, 0).Find(&feedback, self)
	return feedback, totalRecords, err
}

func (self *Feedback) SearchByPage() ([]Feedback, int, error) {
	var totalRecords int
	var feedback []Feedback
	err := orm.Find(&feedback, self)
	totalRecords = len(feedback)
	feedback = feedback[:0]
	err = orm.OrderBy(self.GetSortProperties()[0].Column+" "+self.GetSortProperties()[0].Direction).Limit(self.GetPageSize(), self.GetDisplayStart()).Find(&feedback, self)
	return feedback, totalRecords, err
}
