package controllers

import "github.com/astaxie/beego"

type GroupController struct {
	beego.Controller
}


func (c *GroupController) Get() {
	c.TplName = "file.html"
}
func (c *GroupController) Post() {
	c.TplName = "file.html"
}