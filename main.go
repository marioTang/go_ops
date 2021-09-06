package main

import (
	_ "cmdb/routers"
	"cmdb/utils"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	utils.Regismode()
	//utils.K8s()
	utils.Newk8s()
	beego.Run("0.0.0.0:8080")

}



