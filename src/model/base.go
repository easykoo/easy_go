package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lunny/xorm"
	"github.com/qiniu/log"

	"util"
)

var orm *xorm.Engine

func SetEngine() *xorm.Engine {
	log.Println("db initializing...")
	var err error
	orm, err = xorm.NewEngine("mysql", "root:pass@/easy_go?charset=utf8")
	util.PanicIf(err)
	return orm
}
