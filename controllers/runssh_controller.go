package controllers

import (
	"cmdb/models"
	"cmdb/utils"
)

type RunSshController struct {
	CheckLoginController
}

func (c *RunSshController) Get() {
	c.Data["datas"] = models.SelectSshhost()
	c.TplName = "sshrun.html"
}
func (c *RunSshController) Post() {
	hostname := c.GetString("selectdata")
	cmd := c.GetString("cmd")
	publcikey :=models.SelctSshhostId(hostname).Publcikey
	username :=models.SelctSshhostId(hostname).Username
	port :=models.SelctSshhostId(hostname).Port

    data := utils.Runshell(hostname,username,"",port,publcikey,cmd)

	c.Data["json"] = map[string]interface{}{"code": 2, "message": string(data)}
	c.ServeJSON()

}

