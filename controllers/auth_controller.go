package controllers

type AuthController struct {
	BaseController
}

func (this *AuthController) Login() {
	this.Ctx.WriteString("login post")
}

func (this *AuthController) Logout() {
	this.Ctx.WriteString("logout")
}

func (this *AuthController) Register() {
	this.Ctx.WriteString("register post")
}
