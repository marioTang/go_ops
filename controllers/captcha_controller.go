package controllers

import (
	"encoding/base64"
	"github.com/astaxie/beego"
	"github.com/mojocn/base64Captcha"
	"strings"
)


type CaptchaController struct {
	beego.Controller

}

func (c *CaptchaController) Get() {
	//数字验证码配置
	var store = base64Captcha.DefaultMemStore
	var configD = base64Captcha.DriverDigit{
		Height:   50,
		Width:    300,
		MaxSkew:  0.7,
		DotCount: 80,
		Length:   4,
	}
	base64stringD := base64Captcha.NewCaptcha(&configD, store)
	id, b64s, _ := base64stringD.Generate()
	b64s = strings.Replace(b64s, "data:image/png;base64,", "", 1)
	c.SetSession("captcha", 	store.Get(id, true))
	resp := c.Ctx.ResponseWriter
	resp.Header().Set("Content-Type", "image/png")
	ddd, _ := base64.StdEncoding.DecodeString(b64s)
	resp.Write([]byte(ddd))

	//c.Data["Data"] = resp
	c.TplName = "login.html"
}
func (this *CaptchaController) Post() {

}




