package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"time"
)

type MonitorController struct {
	CheckLoginController
}

func (this *MonitorController) Get() {
	this.TplName = "monitor.html"
}
func (c *MonitorController) Post() {
	groupname := c.GetString("groupname")
	projectname := c.GetString("projectname")
	hostname := c.GetString("hostname")
	scriptpath := c.GetString("scriptpath")
	crontime :=c.GetString("crontime")


	datatime := time.Now()
	var data = models.Monitor{ Groupname : groupname, Hostname: hostname,Projectname: projectname,Scriptpath: scriptpath,Createtime: datatime, Cron: crontime}
	models.InsertMonitor(data)
	c.Data["json"] = map[string]interface{}{"code": 2, "message": "创建定时任务成功"}

	//添加定时任务进程
	utils.Addcron(hostname)
	c.ServeJSON()

}
