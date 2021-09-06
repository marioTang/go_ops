package controllers

import "cmdb/models"

type IndexController struct {
	CheckLoginController
}

func (c *IndexController) Get() {
	loginuser := c.GetSession("loginuser").(string)

	data := models.GetAlibill(loginuser)
	c.Data["Availablecashamount"] = data.Availablecashamount
	c.Data["Availableamount"] = data.Availableamount
	c.TplName = "monitor.html"


}
func (c *IndexController) Post() {
c.TplName = "monitor.html"
}