package utils

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

)


type Ceshi struct {
Id int64
Name string	`orm:"size(64)"`
Passwords string	`orm:"size(64)"`
Baby []*Baby `orm:"reverse(many)"`
}
type Baby struct {
	Id int64
	Name string	`orm:"size(64)"`
	Ceshi *Ceshi `orm:"rel(fk)"`
}



func InitMysql()  {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	driverName := beego.AppConfig.String("driverName")
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"+"&loc=Asia%2FShanghai"

	orm.Debug = true


	//登陆方法
	errs :=orm.RegisterDriver(driverName,orm.DRMySQL)
	if errs != nil {
		fmt.Println("数据驱动注册失败")

	}else {
		fmt.Println("数据驱动注册成功")

		err :=orm.RegisterDataBase("default" ,driverName, dbConn )
		if err != nil {
			fmt.Println("mysql连接失败")
			fmt.Println(err)

		}else {
			fmt.Println("自动创建表结构")
			orm.RegisterModel(new(Ceshi), new(Baby))
			//orm.RegisterModel(new(Userceshi))
			//orm.RegisterModel(new(Post))



			orm.RunSyncdb("default", true, true)

		}

	}


}
