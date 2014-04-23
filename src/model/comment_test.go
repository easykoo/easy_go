package model

import (
	. "common"

	"testing"
)

func init() {
	SetConfig()
	SetLog()
	SetEngine()
}

func Test_GenerateCommentId(t *testing.T) {
//	Init()
	comment := new(Comment)
	comment.Blog.Id = 1
	id, err := comment.GenerateSeq()
	PanicIf(err)

	Expect(t, id, 1)
}
