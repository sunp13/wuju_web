package controllers

import (
	"fmt"
	"strings"
	"wuju_web/entity"
	"wuju_web/models"
	"wuju_web/utils"

	"github.com/beego/beego/v2/server/web"
)

type AsiaController struct {
	web.Controller
}

// JSONList
// @router /asia/list.json [get]
func (c *AsiaController) JSONList() {
	var result entity.JSONResult

	commID := strings.TrimSpace(c.GetString("comm_id"))

	if commID == "" {
		result.Code = 1
		result.Message = "请求参数不正确"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	data, err := models.AsianHandModel.GetListByCommID(commID)
	if err != nil {
		result.Code = 1
		result.Message = "系统异常"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	for i, v := range data {
		homeHandicap := fmt.Sprintf("%v", v["home_handicap"])
		data[i]["home_handicap"] = utils.UtilTranslate(homeHandicap)
	}

	result.Result = data
	result.Message = "成功"
	c.Data["json"] = result
	c.Data["json"] = result
	c.ServeJSON()
}
