package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"github.com/astaxie/beego"
)
type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

type LoginController struct {
	beego.Controller

}

func (c *LoginController) Get() {

	c.TplName = "login.html"
}
func (this *LoginController) Post() {
	captcha := this.GetString("captcha")
	captchasession :=this.GetSession("captcha")
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:", username, ",password:", password)


	//获取用户的用户id
	id := models.GetUserID(username,utils.MD5(password))
	this.SetSession("loginuser", username)

	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	}
	if captcha !=captchasession{
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "验证码错误"}

	}
	if id == 0{
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户或者密码错误"}
	}
	this.ServeJSON()
}




