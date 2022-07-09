package controllers

import (
	"fmt"
	"strings"
	"wuju_web/entity"
	"wuju_web/models"

	"github.com/beego/beego/v2/server/web"
)

type GoalController struct {
	web.Controller
}

// @router /goal/list.json [get]
func (c *GoalController) JSONList() {
	var result entity.JSONResult
	commID := strings.TrimSpace(c.GetString("comm_id"))
	if commID == "" {
		result.Code = 1
		result.Message = "请求参数不正确"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	data, err := models.GoalModel.GetListByCommID(commID)
	if err != nil {
		result.Code = 1
		result.Message = "系统异常"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	for i, v := range data {
		overName := fmt.Sprintf("%v", v["over_name"])
		overName = strings.Replace(overName, ",", "/", -1)
		overName = strings.Replace(overName, " ", "", -1)
		data[i]["over_name"] = overName
	}

	result.Result = data
	c.Data["json"] = result
	c.ServeJSON()
}
