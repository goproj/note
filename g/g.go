package g

import (
	"github.com/astaxie/beego"
)

var (
	Static string
)

func ReadConfig() {
	Static = beego.AppConfig.String("static")
}
