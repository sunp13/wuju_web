package controllers

import (
	"fmt"
	"strconv"
	"strings"
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
	// 开始时间
	beginUnix := time.Now().Add(-10 * time.Minute).Unix()
	// 最大时间
	maxUnix := time.Now().Add(24 * time.Hour).Unix()
	data, err := models.ComingModel.GetList(fmt.Sprintf("%d", beginUnix), fmt.Sprintf("%d", maxUnix))
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

// @router /upcoming/home.json [get]
func (c *UpComingController) JSONGetHomeList() {
	var result entity.JSONResult

	data, err := models.ComingModel.GetHomeList()
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

// @router /upcoming/away.json [get]
func (c *UpComingController) JSONGetAwayList() {
	var result entity.JSONResult

	data, err := models.ComingModel.GetAwayList()
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

// @router /upcoming/history_query.json [post]
func (c *UpComingController) JSONHistoryQuery() {
	var result entity.JSONResult

	local, _ := time.LoadLocation("Asia/Shanghai")

	// 参数
	leagueID := strings.TrimSpace(c.GetString("league_id"))
	beginTime := strings.TrimSpace(c.GetString("begin_time"))
	endTime := strings.TrimSpace(c.GetString("end_time"))
	homeID := strings.TrimSpace(c.GetString("home_id"))
	awayID := strings.TrimSpace(c.GetString("away_id"))

	if beginTime != "" {
		bTime, _ := time.ParseInLocation("2006-01-02 15:04:05", beginTime, local)
		beginTime = fmt.Sprintf("%d", bTime.Unix())
	}

	if endTime != "" {
		eTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, local)
		endTime = fmt.Sprintf("%d", eTime.Unix())
	}

	data, err := models.ComingModel.GetFilterList(leagueID, beginTime, endTime, homeID, awayID)
	if err != nil {
		result.Code = 1
		result.Message = "系统异常"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}

	for i, v := range data {
		commTime := fmt.Sprintf("%v", v["comm_time"])
		commTimeInt64, _ := strconv.ParseInt(commTime, 10, 64)
		commTimer := time.Unix(commTimeInt64, 0)
		data[i]["comm_time"] = commTimer.Format("2006-01-02 15:04:05")
	}

	result.Result = data
	c.Data["json"] = result
	c.ServeJSON()
}
