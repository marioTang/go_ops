package controllers
import (
	"cmdb/models"
	"github.com/astaxie/beego"
)
type FileUploadController struct {
	beego.Controller
}


//func (c *FileUploadController) Get() {
//    	c.TplName = "upload.html"
//
//}
//
//func (c *FileUploadController) Post() {
//	// uploadfilename，这是一个key值，对应的是html中input type-‘file’的name属性值
//	f, h, err := c.GetFile("uploadfilename")
//	if err != nil {
//		log.Fatal("getfile err", err)
//		c.Data["status"] = "上传失败"
//		c.TplName = "upload.html"
//	}else {
//		c.Data["status"] = "上传成功"
//		c.TplName = "upload.html"
//	}
//
//	// 关闭上传的文件，不然的话会出现临时文件不能清除的情况
//	defer f.Close()
//	// 保存位置在 static/upload, 没有文件夹要先创建
//	c.SaveToFile("uploadfilename", "static/upload/" + h.Filename)
//	// html页面
//	c.TplName = "upload.html"
//}


func (c *FileUploadController) Get() {
	c.Data["datas"] = models.SelectMonitor()


	c.TplName = "cronupload.html"
}

func (c *FileUploadController) Post() {

		c.TplName = "cronupload.html"
	}

