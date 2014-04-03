package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"

	. "common"
)

var orm *xorm.Engine

func SetEngine() *xorm.Engine {
	Log.Info("db initializing...")
	var err error
	orm, err = xorm.NewEngine("mysql", "root:pass@/easy_go?charset=utf8")
	PanicIf(err)
	orm.ShowSQL = true
	orm.Logger = Log.GetWriter()
	return orm
}
