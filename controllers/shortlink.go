package controllers

import (
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"net/http"
	"shortlink/helper"
	"shortlink/models"
	"strconv"
	"strings"
	"time"
)

// ShortlinkController operations for Shortlink
type ShortlinkController struct {
	BaseController
}

// URLMapping ...
func (c *ShortlinkController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Shortlink
// @Param	body		body 	models.Shortlink	true		"body for Shortlink content"
// @Success 201 {int} models.Shortlink
// @Failure 403 body is empty
// @router / [post]
func (c *ShortlinkController) Post() {
	var v models.Shortlink
	if parseUserId != 0 {
		v.CreatedBy = parseUserId
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddShortlink(&v); err == nil {
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

// GetOne ...
// @Title Get One
// @Description get Shortlink by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Shortlink
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ShortlinkController) GetOne() {
	slug := c.Ctx.Input.Param(":id")
	v, err := models.GetShortlinkByAlias(slug)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.Data["json"] = helper.JsonResponse(http.StatusNotFound, "Shortlink not found")
	} else {
		if v.Status == 0 {
			c.Ctx.Output.SetStatus(http.StatusNotFound)
			c.Data["json"] = helper.JsonResponse(http.StatusNotFound, "Shortlink is disabled")
		}
		if v.Expire > int(time.Now().Unix()) {
			c.Ctx.Output.SetStatus(http.StatusGone)
			c.Data["json"] = helper.JsonResponse(http.StatusGone, "Shortlink expired")
		} else {
			o := orm.NewOrm()
			v.TotalClick = v.TotalClick + 1
			if _, err = o.Update(v, "total_click"); err != nil {

			}
			v.TotalClick = v.TotalClick + 1
			if v.Password != "" {
				c.Ctx.Output.SetStatus(http.StatusUnauthorized)
				c.Data["json"] = map[string]interface{}{
					"id":       v.Id,
					"longUrl":  v.LongUrl,
					"password": v.Password,
					"aliasUrl": v.AliasUrl,
					"status":   v.Status,
				}
			} else {
				c.Ctx.Output.SetStatus(http.StatusOK)
				c.Data["json"] = map[string]interface{}{
					"id":       v.Id,
					"longUrl":  v.LongUrl,
					"aliasUrl": v.AliasUrl,
					"status":   v.Status,
				}
			}

		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Shortlink
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Shortlink
// @Failure 403
// @router / [get]
func (c *ShortlinkController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	var createdBy string
	if parseUserId != 0 {
		createdBy = strconv.Itoa(parseUserId)
		query["created_by__exact"] = createdBy
		query["status__exact"] = "10"
	} else {
		query["created_by__isnull"] = "true"
		query["status__exact"] = "10"
	}
	l, err := models.GetAllShortlink(
		query,
		fields,
		sortby,
		order,
		offset,
		limit)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusNoContent)
		c.Data["json"] = helper.JsonResponse(http.StatusNoContent, "No Content")
	} else {
		if len(l) > 0 {
			c.Ctx.Output.SetStatus(http.StatusOK)
			c.Data["json"] = l
		} else {
			c.Ctx.Output.SetStatus(http.StatusNoContent)
			c.Data["json"] = helper.JsonResponse(http.StatusNoContent, "No Content")
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Shortlink
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Shortlink	true		"body for Shortlink content"
// @Success 200 {object} models.Shortlink
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ShortlinkController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Shortlink{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateShortlinkById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Shortlink
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ShortlinkController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteShortlink(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
