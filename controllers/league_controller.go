package controllers

import (
	"wuju_web/entity"
	"wuju_web/models"

	"github.com/beego/beego/v2/server/web"
)

type LeagueController struct {
	web.Controller
}

// @router /league/list.json [get]
func (c *LeagueController) JSONGetList() {
	var result entity.JSONResult
	data, err := models.LeagueModel.GetList()
	if err != nil {
		result.Code = 1
		result.Message = "系统异常"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	result.Message = "succ"
	result.Result = data
	c.Data["json"] = result
	c.ServeJSON()
}
