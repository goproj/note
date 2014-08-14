package main

import (
	"github.com/astaxie/beego"
	"github.com/goproj/note/g"
	_ "github.com/goproj/note/routers"
)

func main() {
	g.ReadConfig()
	beego.TemplateLeft = "{["
	beego.TemplateRight = "]}"
	beego.Run()
}
