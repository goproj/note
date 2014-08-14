package controllers

import (
	"github.com/astaxie/beego"
	"github.com/goproj/note/g"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["Static"] = g.Static
}
