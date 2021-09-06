package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"strconv"
)

type UserController struct {
	CheckLoginController}


func (c *UserController) Get() {
	pa :=c.GetString("page","1")
	paint, _ := strconv.Atoi(pa)
	pre_page :=8
	totals,_:=models.CoutUser()
	res :=utils.Paginator(paint, pre_page, totals)
	datalist,_:=models.GetUserPag(pre_page,(paint-1)*pre_page)
	c.Data["datas"] = datalist
	c.Data["paginator"] = res
	c.Data["totals"] = totals
	c.TplName = "user.html"

	editid := c.GetString("edit")

	if len(editid) != 0 {
		paint, _ := strconv.Atoi(editid)
		data := models.GetUserData(paint)
		c.Data["Id"] = data.Id
		c.Data["Name"] = data.Name
		c.Data["Passwd"] = data.Passwd
		c.Data["Projectteam"] = data.Projectteam
		c.TplName = "adduser.html"
	} else {
	}
	delid := c.GetString("del")

	if len(delid) != 0{


		delids, _ := strconv.Atoi(delid)
		models.DelUser(delids)
		data := "删除成功"
		url := "/user"
		c.Data["Data"] = data
		c.Data["Url"] = url
		c.TplName = "status.html"

	}



}

//func (c *UserinfoController) Post() {
//	id := len(c.GetString("editid"))
//	if 	id == 0{
//		username := c.GetString("username")
//		password := c.GetString("password")
//		remarks  := c.GetString("remarks")
//		accesskeyid := c.GetString("accesskeyid")
//		accesskeysecret := c.GetString("accesskeysecret")
//		editro := c.GetString("editro")
//		deldetro := c.GetString("deldetro")
//		groupname :=c.GetString("groupname")
//		if models.SelctUserName(username).Name != username{
//			data :=models.Users{Name: username,Passwd: password,Groupname:groupname,Remarks: remarks,Accesskeysecret: accesskeysecret,Accesskeyid: accesskeyid,Edit: editro,Deldet: deldetro}
//			models.InsertUser(data)
//		}else {
//			c.Data["json"] = map[string]interface{}{"code": 2, "message": "用户已存在"}
//			c.ServeJSON()
//		}
//	}else{
//		fmt.Println("bu wei kong ")
//		paint, _ := strconv.Atoi(c.GetString("editid"))
//
//		username := c.GetString("username")
//		password := c.GetString("password")
//		remarks  := c.GetString("remarks")
//		accesskeyid := c.GetString("accesskeyid")
//		accesskeysecret := c.GetString("accesskeysecret")
//		editro := c.GetString("editro")
//		deldetro := c.GetString("deldetro")
//		groupname :=c.GetString("groupname")
//
//		models.UpdateRoles(username,password,groupname,remarks,editro,deldetro,accesskeyid,accesskeysecret,paint)
//		c.Data["json"] = map[string]interface{}{"code": 2, "message": "更新成功"}
//		c.ServeJSON()
//	}
//	c.Data["json"] = map[string]interface{}{"code": 2, "message": "创建成功"}
//	c.ServeJSON()
//}