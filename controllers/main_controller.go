package controllers

import "github.com/beego/beego/v2/server/web"

type MainController struct {
	web.Controller
}

// @router / [get]
func (c *MainController) PageIndex() {
	c.TplName = "index.html"
}
