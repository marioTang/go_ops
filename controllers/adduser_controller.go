package controllers

type AdduserController struct {
	CheckLoginController}


func (c *AdduserController) Get() {
    	c.TplName = "adduser.html"

}

func (c *AdduserController) Post() {

	//id,_:= strconv.Atoi(c.GetString("editid"))
	//edit:= c.GetString("edit")
	//deldet:=c.GetString("deldet")
	//accesskeyid:=c.GetString("accesskeyid")
	//accesskeysecret := c.GetString("accesskeysecret")
	//name := c.GetString("username")

	//models.UpdateRoles(name,edit,deldet,accesskeyid,accesskeysecret,id)
	data := "更新成功"
	url := "/userinfo"
	c.Data["Data"] = data
	c.Data["Url"] = url
	c.TplName = "status.html"
}