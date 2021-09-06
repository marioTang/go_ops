package controllers

import (
	"cmdb/models"
	"github.com/astaxie/beego"
	"strings"
)

type FileController struct {
	beego.Controller
}
	func (c *FileController) Get() {
		c.Data["datas"] = models.SelectSshhost()

		c.TplName = "file.html"

	}



func (c *FileController) Post() {
	f, h, _ := c.GetFile("image")
	//得到文件的名称
	fileName := h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	f.Close()

	c.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	c.ServeJSON()


}

func (c *FileController) responseErr(err error) {

}