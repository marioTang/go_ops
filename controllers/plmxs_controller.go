package controllers

import (
	"cmdb/models"
	"cmdb/utils"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type PlxmxsController struct {
	CheckLoginController}


func (c *PlxmxsController) Get() {
	options := c.GetString("selectdata")
	fmt.Println(options)

	c.Data["datas"] = models.SelectSshhost()
	c.TplName = "plmxs.html"
}
//func (c *PlxmxsController) Post() {
//	options := c.GetString("options")
//	ue := c.GetString("ue")
//	name := c.GetString("filename")
//
//	now := time.Now()
//	filedir := fmt.Sprintf("static/upload/%s/%d/%d/%d/", "yml", now.Year(), now.Month(), now.Day())
//
//	error  := os.MkdirAll(filedir, os.ModePerm)
//	if error != nil {
//		fmt.Println(error)
//		return
//	}
//
//	timestamp := time.Now().Unix()
//	dqpath,_:=os.Getwd()
//	timestampstr := time.Now().Format("2006-01-02 15:04:05")
//	filename := fmt.Sprintf("%d-%s", timestamp, name + ".yml")
//	filepathstr := filepath.Join(filedir, filename)
//
//	f,err := os.Create(filedir + filename)
//	defer f.Close()
//	if err !=nil {
//		fmt.Println(err.Error())
//	} else {
//		_,err=f.Write([]byte(ue))
//		fmt.Println(err)
//	}
//
//	data := models.Upload{0,options,filename,name +".yml",filepathstr,"/home",timestampstr}
//    models.PlmxsUpload(data)
//
//	publcikey :=models.SelctSshhostId(options).Publcikey
//	username :=models.SelctSshhostId(options).Username
//	port :=models.SelctSshhostId(options).Port
//
//    os.Chdir(filedir)
//	sourcefile := name +".yml"
//
//	utils.Uploadfile(options,username,"",port,publcikey,"./"+filename,"/home/")
//	utils.Runshell(options,username,"",port,publcikey, "cd /home&&./bash.go " + filename + sourcefile)
//
//	os.Chdir(dqpath)
//	c.TplName = "plmxs.html"
//}

func (c *PlxmxsController) Post() {
	options := c.GetString("options")
	ue := c.GetString("ue")
	name := c.GetString("filename")
	sourcefile := name +".yml"

	now := time.Now()
	filedir := fmt.Sprintf("static/upload/%s/%d/%d/%d/", "yml", now.Year(), now.Month(), now.Day())

	error  := os.MkdirAll(filedir, os.ModePerm)
	if error != nil {
		fmt.Println(error)
		return
	}

	//timestamp := time.Now().Unix()
	dqpath,_:=os.Getwd()
	timestampstr := time.Now().Format("2006-01-02 15:04:05")
	//filename := fmt.Sprintf("%d-%s", timestamp, name + ".yml")
	filepathstr := filepath.Join(filedir, sourcefile)

	f,err := os.Create(filedir + sourcefile)
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		_,err=f.Write([]byte(ue))
		fmt.Println(err)
	}

	data := models.Upload{0,options,"",sourcefile,filepathstr,"/home",timestampstr}
	models.PlmxsUpload(data)

	publcikey :=models.SelctSshhostId(options).Publcikey
	username :=models.SelctSshhostId(options).Username
	port :=models.SelctSshhostId(options).Port

	os.Chdir(filedir)

	utils.Uploadfile(options,username,"",port,publcikey,"./"+sourcefile,"/home/")
	//utils.Runshell(options,username,"",port,publcikey, "cd /home&&./bash.go " + filename + sourcefile)

	os.Chdir(dqpath)
	c.Data["json"] = map[string]interface{}{"code": 1, "message": "创建成功"}
	c.ServeJSON()
}