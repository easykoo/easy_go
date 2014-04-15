package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	. "common"
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
	orm.ShowSQL = Cfg.MustBool("db", "show_sql", false)
	orm.Logger = Log.GetWriter()
	return orm
}
