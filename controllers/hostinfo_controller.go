package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"strconv"
)

type HostinfoController struct {
	CheckLoginController
}


func (c *HostinfoController) Get() {
    pa :=c.GetString("page","1")
	paint, _ := strconv.Atoi(pa)
	pre_page :=15
	totals,_:=models.CoutSshhost()
	res :=utils.Paginator(paint, pre_page, totals)
	datalist,_:=models.GetSshhostPag(pre_page,(paint-1)*pre_page)
	c.Data["datas"] = datalist
	c.Data["paginator"] = res
	c.Data["totals"] = totals
	c.TplName = "sshinfo.html"

	editid := c.GetString("edit")

	if len(editid) != 0 {
		paint, _ := strconv.Atoi(editid)
		data := models.SelctSshhost(paint)
		c.Data["Id"] = data.Id
		c.Data["Hostname"] = data.Hostname
		c.Data["Username"] = data.Username
		c.Data["Password"] = data.Password
		c.Data["Port"] = data.Port
		c.Data["Publcikey"] = data.Publcikey

		c.TplName = "edithost.html"
		} else {
			}
	delid := c.GetString("del")

	if len(delid) != 0{
		delids, _ := strconv.Atoi(delid)
			models.DelSshhost(delids)
			data := "删除成功"
			url := "/sshinfo"
			c.Data["Data"] = data
			c.Data["Url"] = url
			c.TplName = "status.html"
		}
	}



func (c *HostinfoController) Post() {
	     fmt.Println("ffffffff")
		id,_:= strconv.Atoi(c.GetString("editid"))
		hostname:= c.GetString("hostname")
		username:=c.GetString("username")
		password:=c.GetString("password")
		port,_:= strconv.Atoi(c.GetString("port"))
		publcikey := c.GetString("publcikey")

		models.UpdateSshhost(hostname,username,password,port,publcikey,id)
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "数据更新成功"}
	c.ServeJSON()

}