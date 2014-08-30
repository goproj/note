package routers

import (
	"github.com/astaxie/beego"
	"github.com/goproj/note/controllers"
)

func init() {
	beego.Router("/", &controllers.AuthController{}, "*:Index")
	beego.Router("/health", &controllers.HealthController{})
	beego.Router("/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/register", &controllers.AuthController{}, "post:Register")
	beego.Router("/logout", &controllers.AuthController{}, "*:Logout")
	beego.Router("/me", &controllers.MainController{}, "get:Me")

	beego.Router("/note/todo", &controllers.NoteController{}, "get:Todo")
	beego.Router("/note/history", &controllers.NoteController{}, "get:History")
	beego.Router("/note", &controllers.NoteController{}, "post:Add")
	beego.Router("/note/:id/mark/:done", &controllers.NoteController{}, "get:Mark")
	beego.Router("/note/:id", &controllers.NoteController{}, "get:Read")
	beego.Router("/note/:id", &controllers.NoteController{}, "delete:Delete")
	beego.Router("/note/:id", &controllers.NoteController{}, "put:Update")
}
