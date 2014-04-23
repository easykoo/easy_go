package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	. "common"
	"time"
)

var orm *xorm.Engine

func SetEngine() *xorm.Engine {
	Log.Info("db initializing...")
	var err error
	username := Cfg.MustValue("db", "username", "root")
	password := Cfg.MustValue("db", "password", "pass")
	dbName := Cfg.MustValue("db", "db_name", "easy_go")
	orm, err = xorm.NewEngine("mysql", username+":"+password+"@/"+dbName+"?charset=utf8")
	PanicIf(err)
	orm.TZLocation = time.Local
	orm.ShowSQL = Cfg.MustBool("db", "show_sql", false)
	orm.Logger = Log
	return orm
}

type DbUtil struct {}

func (self *DbUtil) GetRecentComments() (comments []Comment) {
	err := orm.OrderBy("create_date desc").Limit(5, 0).Find(&comments, &Comment{})
	PanicIf(err)
	return
}
