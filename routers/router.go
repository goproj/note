package routers

import (
	"github.com/astaxie/beego"
	"github.com/goproj/note/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthController{})
	beego.Router("/login", &controllers.AuthController{}, "*:Login")
	beego.Router("/register", &controllers.AuthController{}, "*:Register")
	beego.Router("/logout", &controllers.AuthController{}, "*:Logout")
	beego.Router("/me", &controllers.MainController{}, "get:Me")
}
