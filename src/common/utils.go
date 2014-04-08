package common

import (
	"io"
	"os"
	"strconv"
)

var Log *Logger

func init() {
	var w io.Writer
	if Cfg.MustBool("", "log_file") {
		f, _ := os.OpenFile(Cfg.MustValue("", "log_path"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
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
