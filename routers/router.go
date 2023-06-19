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
			beego.NSRouter("/info", &controllers.AuthController{}, "get:GetOne"),
			beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
			beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
		),
		beego.NSNamespace("/shortlink",
			beego.NSInclude(&controllers.ShortlinkController{}),
			beego.NSRouter("/", &controllers.ShortlinkController{}, "get:GetAll"),
			beego.NSRouter("/?:id", &controllers.ShortlinkController{}, "put:Put"),
			beego.NSRouter("/manager/?:id", &controllers.ShortlinkController{}, "get:GetManager"),
			beego.NSRouter("/?:id/?:password", &controllers.ShortlinkController{}, "get:GetOne"),
			beego.NSRouter("/", &controllers.ShortlinkController{}, "post:Post"),
		),
	)
	beego.ErrorHandler("404", NotFoundHandler)
	beego.ErrorHandler("403", ForbiddenHandler)
	beego.ErrorHandler("500", InternalServerErrorHandler)
	beego.AddNamespace(ns)
}
