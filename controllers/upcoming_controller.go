package controllers

import (
	"fmt"
	"strconv"
	"time"
	"wuju_web/entity"
	"wuju_web/models"

	"github.com/beego/beego/v2/server/web"
)

type UpComingController struct {
	web.Controller
}

// 获取即将开赛的列表 按时间倒叙
// @router /upcoming/list.json [get]
func (c *UpComingController) JSONList() {
	var result entity.JSONResult
	// 当前时间
	maxUnix := time.Now().Add(6 * time.Minute).Unix()
	data, err := models.ComingModel.GetList(fmt.Sprintf("%d", maxUnix))
	if err != nil {
		result.Code = 1
		result.Message = "系统异常"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	for i, v := range data {
		vUnix := fmt.Sprintf("%v", v["comm_time"])
		vUnixInt64, _ := strconv.ParseInt(vUnix, 10, 64)
		vTime := time.Unix(vUnixInt64, 0)
		data[i]["comm_time"] = vTime.Format("2006-01-02 15:04:05")
	}

	result.Result = data
	c.Data["json"] = result
	c.ServeJSON()
}
