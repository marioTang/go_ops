package controllers

import (
	"cmdb/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}
type Host struct {
	Cpu_model string
	Cpu_num string
	Memory string
	Disk string
	Intranet_ip string
	Hostname string
	Os string
}


func (c *ApiController) Get() {

}

func (c *ApiController) Post() {
	//新增逻辑
	var ob Host
	body := c.Ctx.Input.RequestBody

	json.Unmarshal(body, &ob)

	selectdata := models.SelectAsset(ob.Intranet_ip)
	fmt.Println(selectdata)

	if ob.Intranet_ip != selectdata.Intranet_ip {
		data := models.Asset{Cpu_model: ob.Cpu_model, Cpu_num: ob.Cpu_num, Intranet_ip: ob.Intranet_ip, Os: ob.Os, Disk: ob.Disk, Hostname: ob.Hostname, Memory: ob.Memory}
		models.InsertAsset(data)

	}else {
		if (ob.Cpu_model != selectdata.Cpu_model || ob.Hostname != selectdata.Hostname || ob.Cpu_num != selectdata.Cpu_num || ob.Memory != selectdata.Memory || ob.Os != selectdata.Os || ob.Disk != selectdata.Disk) {
			models.UpdateAsset(ob.Intranet_ip, ob.Cpu_model, ob.Hostname, ob.Os, ob.Cpu_num, ob.Memory, ob.Disk)
		}

		c.TplName = "asset.html"
	}
	c.TplName = "asset.html"

}