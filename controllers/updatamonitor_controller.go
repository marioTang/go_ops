package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"strconv"
)

type Updatamonitor struct {
CheckLoginController
}

func (c *Updatamonitor) Get() {

	//	//获取参数
//	executionid := c.GetString("execution")
//	delid := c.GetString("del")
//	editid := c.GetString("edit")
//
//	//分页功能
//	pa :=c.GetString("page","1")
//	paint, _ := strconv.Atoi(pa)
//	pre_page :=15
//	totals,_:=models.CoutMonitor()
//	res :=utils.Paginator(paint, pre_page, totals)
//	datalist,_:=models.GetMonitorPag(pre_page,(paint-1)*pre_page)
//	c.Data["datas"] = datalist
//	c.Data["paginator"] = res
//	c.Data["totals"] = totals
//	c.TplName = "monitorinfo.html"
//
//	//删除功能
//	if len(delid) != 0{
//
//		//删除定时任务
//		taskname := models.SelctMonitorHost(delid).Projectname
//		fmt.Println(taskname)
//		utils.Delcron(taskname)
//
//		//删除定时任务主机
//		delids, _ := strconv.Atoi(delid)
//		models.DelMonitor(delids)
//
//		data := "删除成功"
//		url := "/monitorinfo"
//		c.Data["Data"] = data
//		c.Data["Url"] = url
//		c.TplName = "status.html"
//	}
//
//	//立即执行模块
//	if len(executionid) != 0{
//
//		//host := term["hostname"] .(string)
//		host := models.SelctMonitorHost(executionid).Hostname
//		scriptpath := models.SelctMonitorHost(executionid).Scriptpath
//		publcikey :=models.SelctSshhostId("119.3.106.91").Publcikey
//		runshell := string(utils.Runshell(host, "root", "", 22, publcikey, scriptpath))
//		c.Data["Runshell"] = runshell
//		c.TplName = "execdisplay.html"
//
//	}
//
//	if len(editid) != 0 {
//		data := models.SelctMonitorHost(editid)
//		c.Data["Id"] = data.Id
//		c.Data["Groupname"] = data.Groupname
//		c.Data["Projectname"] = data.Projectname
//		c.Data["Hostname"] = data.Hostname
//		c.Data["Scriptpath"] = data.Scriptpath
//		c.Data["Cron"] = data.Cron
//		c.TplName = "editmonitor.html"
//	}
//
//
}
func (c *Updatamonitor) Post() {
	id,_:= strconv.Atoi(c.GetString("editid"))
	groupname := c.GetString("groupname")
	hostname :=c.GetString("hostname")
	projectname :=c.GetString("projectname")
	scriptpath := c.GetString("scriptpath")
	crontime := c.GetString("crontime")

	models.UpdateMonitor(groupname,hostname,projectname,scriptpath,crontime,id)

	//更新定时任务(先删除后添加)
	taskname := models.SelctMonitorId(hostname).Projectname
	utils.Delcron(taskname)
	utils.Addcron(hostname)
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "定时任务更新成功"}
	c.ServeJSON()

}



