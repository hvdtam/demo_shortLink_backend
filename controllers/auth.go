package controllers

import (
	"encoding/json"
	"fmt"
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
	c.Mapping("GetOne", c.GetOne)
}

func (c *AuthController) GetOne() {
	userSession := c.GetSession("current_user")
	if userSession != nil {
		v, err := models.GetUserById(userSession.(map[string]int)["id"])
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			resp := make(map[string]interface{})
			resp["status"] = http.StatusOK
			resp["user"] = v
			resp["code"] = http.StatusOK
			c.Data["json"] = resp
		}
	} else {
		c.Data["json"] = helper.JsonResponse(http.StatusUnauthorized, "Unauthorized")
	}
	c.ServeJSON()
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
			accessToken, _ := helper.GenerateToken(v.Username, v.Password)
			globalAccessToken = accessToken
			userSession := map[string]int{"id": v.Id}
			err := c.SetSession("current_user", userSession)
			if err != nil {
				fmt.Println(err)
			}
			c.Data["json"] = helper.AccessToken(http.StatusCreated, accessToken)
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
			accessToken, _ := helper.GenerateToken(v.Username, v.Password)
			c.Data["json"] = helper.AccessToken(200, accessToken)
			o := orm.NewOrm()
			user := &models.Users{Username: v.Username}
			_ = o.Read(user, "username")
			globalAccessToken = accessToken
			userSession := map[string]int{"id": user.Id}
			err := c.SetSession("current_user", userSession)
			if err != nil {
				fmt.Print(err)
			}
		} else {
			c.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			c.Data["json"] = helper.JsonResponse(http.StatusUnauthorized, err.Error())
		}
	} else {
		c.Data["json"] = helper.JsonResponse(http.StatusBadRequest, err.Error())
	}
	c.ServeJSON()
}
