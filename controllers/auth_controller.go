package controllers

type AuthController struct {
	BaseController
}

func (this *AuthController) LoginGet() {
	this.Ctx.WriteString("login get")
}

func (this *AuthController) LoginPost() {
	this.Ctx.WriteString("login post")
}

func (this *AuthController) Logout() {
	this.Ctx.WriteString("logout")
}

func (this *AuthController) RegisterGet() {
	this.Ctx.WriteString("register get")
}

func (this *AuthController) RegisterPost() {
	this.Ctx.WriteString("register post")
}
