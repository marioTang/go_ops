package controllers

import (
	"github.com/astaxie/beego"
)


type TestController  struct {
	beego.Controller

}

func (c *TestController) Test() {
	c.Ctx.WriteString(`我是test控制器中的test方法`)
}

func (c *TestController) Testing() {
	c.Ctx.WriteString(`我是test控制器中的post方法`)
}

