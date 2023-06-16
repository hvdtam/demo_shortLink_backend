package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"shortlink/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.AuthController{}),
			beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
			beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
		),
	)
	beego.ErrorHandler("404", NotFoundHandler)
	beego.ErrorHandler("403", ForbiddenHandler)
	beego.ErrorHandler("500", InternalServerErrorHandler)
	beego.AddNamespace(ns)
}
