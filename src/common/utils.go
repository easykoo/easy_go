package common

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"reflect"
)

var Log *Logger

func SetLog() {
	var w io.Writer
	if Cfg.MustBool("", "log_file", false) {
		f, _ := os.OpenFile(Cfg.MustValue("", "log_path", "./log.txt"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
		w = io.MultiWriter(f)
	} else {
		w = os.Stdout
	}
	Log = New(w, "", Lshortfile|Ldate|Lmicroseconds)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseInt(value string) int {
	if value == "" {
		return 0
	}
	val, _ := strconv.Atoi(value)
	return val
}

func IntString(value int) string {
	return strconv.Itoa(value)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Atoa(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		c := rune(str[i])
		if 'A' <= c && c <= 'Z' && i > 0 {
			result = result + "_" + strings.ToLower(string(str[i]))
		} else {
			result = result + string(str[i])
		}
	}
	return result
}

/* Test Helpers */
func Expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
