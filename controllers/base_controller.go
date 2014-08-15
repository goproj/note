package controllers

import (
	"github.com/astaxie/beego"
	"github.com/goproj/note/g"
	"github.com/goproj/note/models"
)

type Checker interface {
	CheckLogin()
}

type BaseController struct {
	beego.Controller
	CurrentUser *models.User
}

func (this *BaseController) Prepare() {
	if app, ok := this.AppController.(Checker); ok {
		app.CheckLogin()
	}
	this.Data["Static"] = g.Static
}

func (this *BaseController) ServeErrMsg(msg string) {
	this.Data["json"] = map[string]interface{}{"err": 1, "msg": msg}
	this.ServeJson()
}

func (this *BaseController) ServeOkMsg(msg string) {
	this.Data["json"] = map[string]interface{}{"err": 0, "msg": msg}
	this.ServeJson()
}

func (this *BaseController) ServeOk() {
	this.Data["json"] = map[string]interface{}{"err": 0}
	this.ServeJson()
}

func (this *BaseController) ServeDataJson(val interface{}) {
	this.Data["json"] = map[string]interface{}{"err": 0, "data": val}
	this.ServeJson()
}
