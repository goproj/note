package routers

import (
	"github.com/astaxie/beego"
	"github.com/goproj/note/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthController{})
	beego.Router("/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/register", &controllers.AuthController{}, "post:Register")
	beego.Router("/logout", &controllers.AuthController{}, "*:Logout")
}
