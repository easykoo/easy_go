package model

import (
	. "common"

	"reflect"
	"testing"
)

func Test_user(t *testing.T) {
	SetEngine()
	user := &User{Username: "test4", Password: "11111", Email: "ddd3@ddd.com"}
	user.Delete()
	err := user.Insert()
	PanicIf(err)

	dbUser, err1 := user.GetUser()
	PanicIf(err1)
	expect(t, dbUser.DeptId, 1)
	expect(t, dbUser.RoleId, 3)
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
