package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"strconv"
)

type SeachController struct {
	CheckLoginController
}


func (c *SeachController) Get() {
	//获取参数
	executionid := c.GetString("execution")
	delid := c.GetString("del")
	editid := c.GetString("edit")
	groupname := c.GetString("name")

		pa := c.GetString("page", "1")
		paint, _ := strconv.Atoi(pa)
		pre_page := 15
		totals := models.SeachMonitorCount(groupname)
		res := utils.Paginator(paint, pre_page, totals)
		datalist, _ := models.SeachMonitorData(groupname, pre_page, (paint-1)*pre_page)

	c.Data["datas"] = datalist
	c.Data["name"] = groupname
	c.Data["paginator"] = res
	c.Data["totals"] = totals
	c.TplName = "monitorinfo.html"

	if len(delid) != 0{
		//删除定时任务
		taskname := models.SelctMonitorHost(delid).Projectname
		utils.Delcron(taskname)

		//删除定时任务主机
		delids, _ := strconv.Atoi(delid)
		models.DelMonitor(delids)

		data := "删除成功"
		url := "/monitorinfo"
		c.Data["Data"] = data
		c.Data["Url"] = url
		c.TplName = "status.html"
	}

	//立即执行模块
	if len(executionid) != 0{

		//host := term["hostname"] .(string)
		host := models.SelctMonitorHost(executionid).Hostname
		scriptpath := models.SelctMonitorHost(executionid).Scriptpath
		publcikey :=models.SelctSshhostId("119.3.106.91").Publcikey
		runshell := string(utils.Runshell(host, "root", "", 22, publcikey, scriptpath))
		c.Data["Runshell"] = runshell
		c.TplName = "execdisplay.html"

	}

	if len(editid) != 0 {
		data := models.SelctMonitorHost(editid)
		c.Data["Id"] = data.Id
		c.Data["Groupname"] = data.Groupname
		c.Data["Projectname"] = data.Projectname
		c.Data["Hostname"] = data.Hostname
		c.Data["Scriptpath"] = data.Scriptpath
		c.Data["Cron"] = data.Cron
		c.TplName = "editmonitor.html"
	}
	}

func (c *SeachController) Post() {
	groupname := c.GetString("searchname")

	if len(groupname) != 0 {
		fmt.Println(groupname)
		pa := c.GetString("page", "1")
		paint, _ := strconv.Atoi(pa)
		pre_page := 15
		totals := models.SeachMonitorCount(groupname)
		res := utils.Paginator(paint, pre_page, totals)
		datalist, _ := models.SeachMonitorData(groupname, pre_page, (paint-1)*pre_page)

		fmt.Println(datalist)
		c.Data["datas"] = datalist
		c.Data["paginator"] = res
		c.Data["totals"] = totals
		c.TplName = "monitorinfo.html"
	}
}