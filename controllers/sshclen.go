package controllers

import "cmdb/models"

type SshlistController struct {
	CheckLoginController}


func (c *SshlistController) Get() {

	hosts := models.Getsshws_array_all()
	c.Data["Hosts"] = hosts
	c.TplName = "wsssh.html"

}

func (c *SshlistController) Post() {

}
