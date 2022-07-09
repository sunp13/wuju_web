package controllers

import (
	"strings"
	"wuju_web/entity"
	"wuju_web/models"

	"github.com/beego/beego/v2/server/web"
)

type FulltimeController struct {
	web.Controller
}

// @router /fulltime/list.json [get]
func (c *FulltimeController) JSONList() {
	var result entity.JSONResult
	commID := strings.TrimSpace(c.GetString("comm_id"))
	if commID == "" {
		result.Code = 1
		result.Message = "请求参数不正确"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	data, err := models.FulltimeModel.GetListByCommID(commID)
	if err != nil {
		result.Code = 1
		result.Message = "系统异常"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	result.Result = data
	c.Data["json"] = result
	c.ServeJSON()
}
