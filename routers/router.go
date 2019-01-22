package routers

import (
	"beego-blog/controllers"
	"beego-blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	ns :=
		beego.NewNamespace("/admin",

			beego.NSRouter("/", &admin.IndexController{}, "get:Index"),
			beego.NSRouter("/login", &admin.IndexController{}, "*:Login"),

			beego.NSRouter("/dashboard", &admin.DashboardController{}, "get:Dashboard"),
			beego.NSNamespace("/user",
				beego.NSRouter("/", &admin.UserController{}, "get:UserList"),
				beego.NSRouter("/add", &admin.UserController{}, "get:UserAdd"),
				beego.NSRouter("/add", &admin.UserController{}, "post:Add"),
				beego.NSRouter("/edit/:id", &admin.UserController{}, "get:EditUser"),
		),



		)


	beego.AddNamespace(ns)
}
