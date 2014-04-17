package model

import (
	"fmt"
	"testing"
)

func Test_Category(t *testing.T) {
	SetEngine()
	err := orm.DropTables(&Category{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = orm.CreateTables(&Category{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = orm.Insert(&Category{Id: 1, Description: "test1"}, &Category{Id: 2, Description: "test2"}, &Category{Id: 3, Description: "test3"})
	if err != nil {
		fmt.Println(err)
		return
	}

	category := Category{}
	_, err = orm.Id(1).Get(&category)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(category)

	Expect(t, category.Id, 1)
}
