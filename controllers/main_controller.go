package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "ulric.qin@gmail.com"
	this.TplNames = "index.html"
}
