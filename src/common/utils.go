package common

import (
	"io"
	"os"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

var Log *Logger

func init() {
	f, _ := os.OpenFile(Cfg.MustValue("", "log_path"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	w := io.MultiWriter(f)
	Log = New(w, "", Lshortfile|Ldate|Lmicroseconds)
}
