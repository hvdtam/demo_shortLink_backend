package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
	"shortlink/helper"
	"shortlink/models"
)

// AuthController operations for User
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Register)
	c.Mapping("Post", c.Login)
}

// Register ...
// @Title Register
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Register() {
	var v models.Users
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUser(&v); err == nil {
			access_token, _ := helper.GenerateToken(v.Username, v.PasswordHash)
			userSession := map[string]int{"id": v.Id}
			c.SetSession("current_user", userSession)
			c.Data["json"] = helper.AccessToken(http.StatusCreated, access_token)
		} else {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			c.Data["json"] = helper.JsonResponse(http.StatusBadRequest, err.Error())
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Login ...
// @Title Login
// @Description login User
// @Param	body		body 	models.LoginForm	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Login() {
	var v models.LoginForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.Login(&v); err == nil {
			access_token, _ := helper.GenerateToken(v.Username, v.Password)
			c.Data["json"] = helper.AccessToken(200, access_token)
			o := orm.NewOrm()
			user := &models.Users{Username: v.Username}
			_ = o.Read(user, "username")

			userSession := map[string]int{"id": user.Id}
			c.SetSession("current_user", userSession)
		} else {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			c.Data["json"] = helper.JsonResponse(http.StatusUnauthorized, err.Error())
		}
	} else {
		c.Data["json"] = helper.JsonResponse(http.StatusBadRequest, err.Error())
	}
	c.ServeJSON()
}
