package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"shortlink/helper"
	"shortlink/models"
	"strings"
)

var globalAccessToken string
var parseUserId int

type BaseController struct {
	beego.Controller
}

func (c *ShortlinkController) Prepare() {
	var code int

	code = http.StatusOK
	message := "Invalid Token"
	token := c.Ctx.Request.Header["Authorization"]
	if len(token) == 0 {
		code = http.StatusOK
	} else {

		user, err := helper.ParseToken(strings.TrimPrefix(token[0], "Bearer "))
		if err != nil {
			log.Println(err)
			return
		}

		if user == nil {
			code = http.StatusUnauthorized
			return
		}

		v := &models.Users{Username: user.Username}
		o := orm.NewOrm()
		if err := o.Read(v, "username"); err != nil {

		}
		parseUserId = v.Id
		if err != nil {
			switch err.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				code = http.StatusGone

			default:
				code = http.StatusUnauthorized

			}
		}
	}

	if code != http.StatusOK {
		c.Ctx.Output.SetStatus(code)
		c.Data["json"] = helper.JsonResponse(code, message)
		c.ServeJSON()
	}
}
