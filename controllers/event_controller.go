package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	"wuju_web/entity"
	"wuju_web/models"
	"wuju_web/utils"

	"github.com/beego/beego/v2/server/web"
)

type EventController struct {
	web.Controller
}

// @router /event/list.json [get]
func (c *EventController) JSONList() {
	var result entity.JSONResult
	commID := strings.TrimSpace(c.GetString("comm_id"))
	if commID == "" {
		result.Code = 1
		result.Message = "请求参数不正确"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	data, err := models.EventModel.GetListByCommID(commID)
	if err != nil {
		result.Code = 1
		result.Message = "系统异常"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	for i, v := range data {
		homeDataText := fmt.Sprintf("%v", v["home_data"])
		var homeDataList []map[string]string
		err := json.Unmarshal([]byte(homeDataText), &homeDataList)
		if err != nil {
			continue
		}
		for x, y := range homeDataList {
			HD := y["HD"]
			homeDataList[x]["HD"] = utils.UtilTranslate(HD)
		}
		data[i]["home_data_list"] = homeDataList
		delete(data[i], "home_data")
	}

	for i, v := range data {
		awayDataText := fmt.Sprintf("%v", v["away_data"])
		var dataList []map[string]string
		err := json.Unmarshal([]byte(awayDataText), &dataList)
		if err != nil {
			continue
		}
		for x, y := range dataList {
			HD := y["HD"]
			dataList[x]["HD"] = utils.UtilTranslate(HD)
		}
		data[i]["away_data_list"] = dataList
		delete(data[i], "away_data")
	}

	result.Result = data
	c.Data["json"] = result
	c.ServeJSON()
}
