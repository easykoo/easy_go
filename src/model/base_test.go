package model

import (
	. "common"

	"testing"
)

func Init() {
	SetConfig()
	SetLog()
	SetEngine()
}

func Test_GetHotBlog(t *testing.T) {
	Init()
	blog := new(DbUtil).GetHotBlog()
	Log.Debug(blog)
}
