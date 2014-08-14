package controllers

type HealthController struct {
	BaseController
}

func (this *HealthController) Get() {
	this.Ctx.WriteString("ok")
}
