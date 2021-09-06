package controllers

import (
	"cmdb/models"
	"fmt"
	"strconv"
)

type SshController struct {
	CheckLoginController
}

func (this *SshController) Get() {
	this.TplName = "ssh.html"
}
func (c *SshController) Post() {
	fmt.Println("dddddddddddddd")
	hostname := c.GetString("hostname")
	username := c.GetString("username")
	password := c.GetString("password")
	publcikey := c.GetString("publcikey")
	port := c.GetString("port")
	portint, _ := strconv.Atoi(port)

	if models.SelctSshhostId(hostname).Hostname != ""{
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "主机数据存在"}
		c.ServeJSON()


	}else {
			data := models.Sshhost{Hostname: hostname,Username: username,Password: password,Port:portint, Publcikey:publcikey}
			models.InsertSshhost(data)
		    c.Data["json"] = map[string]interface{}{"code": 2, "message": "主机数据添加成功"}
		    c.ServeJSON()

	}
	}

