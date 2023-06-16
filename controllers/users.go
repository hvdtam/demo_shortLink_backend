package controllers

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"shortlink/models"
)

// UsersController operations for Users
type UsersController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Users
// @Param	body		body 	models.Users	true		"body for Users content"
// @Success 201 {int} models.Users
// @Failure 403 body is empty
// @router / [post]
func (c *UsersController) Post() {
	var v models.Users
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
