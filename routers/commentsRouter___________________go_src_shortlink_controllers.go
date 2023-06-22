package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["shortlink/controllers:AuthController"] = append(beego.GlobalControllerRouter["shortlink/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:AuthController"] = append(beego.GlobalControllerRouter["shortlink/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"] = append(beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"] = append(beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"] = append(beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"] = append(beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"],
		beego.ControllerComments{
			Method:           "GetManager",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"] = append(beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"] = append(beego.GlobalControllerRouter["shortlink/controllers:ShortlinkController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["shortlink/controllers:UsersController"] = append(beego.GlobalControllerRouter["shortlink/controllers:UsersController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
