package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"shortlink/helper"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["json"] = helper.JsonResponse(http.StatusOK, "OK")
	c.ServeJSON()
}
