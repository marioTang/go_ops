package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"strconv"
)

type PlxmxsVersionController struct {
	CheckLoginController
}


func (c *PlxmxsVersionController) Get() {
	pa :=c.GetString("page","1")
	paint, _ := strconv.Atoi(pa)
	pre_page :=15
	totals,_:=models.CoutPlmxs()
	res :=utils.Paginator(paint, pre_page, totals)
	datalist,_:=models.GetUploadPag(pre_page,(paint-1)*pre_page)
	c.Data["datas"] = datalist
	c.Data["paginator"] = res
	c.Data["totals"] = totals
	c.TplName = "plmxsversion.html"

	delid := c.GetString("del")
	rollback := c.GetString("rollback")

	if len(delid) != 0{
		delids, _ := strconv.Atoi(delid)
		models.DelUpload(delids)
		data := "删除成功"
		url := "/plmxsversion"
		c.Data["Data"] = data
		c.Data["Url"] = url
		c.TplName = "status.html"
	}
	if len(rollback) != 0{
		//回滚
		rollback, _ := strconv.Atoi(rollback)
		host := models.SelctUpload(rollback).Host


		sourcefile := models.SelctUpload(rollback).Sourcefile
		filename := models.SelctUpload(rollback).Filename

		pkey :=models.SelctSshhostId(host).Publcikey
		user :=models.SelctSshhostId(host).Username

		//psswd :=models.SelctSshhostId(host).Password
		port :=models.SelctSshhostId(host).Port

		response := fmt.Sprintf("cd /home/&& rm -rf  %d.yml &&mv  %d %d", sourcefile ,filename,sourcefile )

		fmt.Println(string(utils.Runshell(host,user,"",port,pkey, response)))

		data := "回滚成功"
		url := "/plmxsversion"
		c.Data["Data"] = data
		c.Data["Url"] = url
		c.TplName = "status.html"
	}
}



func (c *PlxmxsVersionController) Post() {
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