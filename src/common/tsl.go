package common

import (
	cfg "github.com/Unknwon/goconfig"
	"strings"
	"fmt"
)

var Tsl *cfg.ConfigFile

func init() {
	Tsl, _ = cfg.LoadConfigFile("messages.ini")
}

func Translate(lang string, format string) string {
	if lang == "" || !strings.Contains(lang, "zh") {
		lang = "en"
	}
	return Tsl.MustValue(lang, format)
}

func Translatef(lang string, format string, args ...interface{}) string {
	if lang == "" || !strings.Contains(lang, "zh") {
		lang = "en"
	}
	return fmt.Sprintf(Tsl.MustValue(lang, format), args)
}
