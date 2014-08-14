package routers

import (
	"github.com/astaxie/beego"
	"github.com/goproj/note/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/health", &controllers.HealthController{})
	beego.Router("/login", &controllers.AuthController{}, "get:LoginGet;post:LoginPost")
	beego.Router("/register", &controllers.AuthController{}, "get:RegisterGet;post:RegisterPost")
	beego.Router("/logout", &controllers.AuthController{}, "*:Logout")
}
