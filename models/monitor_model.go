package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Monitor struct {
	Id int
	Groupname     string
	Projectname   string
	Hostname      string
	Scriptpath    string
	Createtime time.Time
	Cron string
}

func InsertMonitor(monitor Monitor ) error {
	o :=orm.NewOrm()
	u := Monitor{0,monitor.Groupname,monitor.Projectname,monitor.Hostname,monitor.Scriptpath,monitor.Createtime, monitor.Cron}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

//
func SelctMonitorId(hostname string) Monitor {
	var u Monitor
	o :=orm.NewOrm()
	err := o.Raw("select * from monitor where hostname = ?",hostname).QueryRow(&u)
	if err != nil {
	}
	return u
}

func SelctMonitorHost(id string) Monitor {
	var u Monitor
	o :=orm.NewOrm()
	err := o.Raw("select * from monitor where id = ?",id).QueryRow(&u)
	if err != nil {
	}
	return u
}


func SelectMonitor() ([]Monitor) {
	o := orm.NewOrm()
	var u []Monitor
	o.QueryTable("monitor").All(&u)
	return u
}

func CoutMonitor() (int64, error) {
	o := orm.NewOrm()
	qs, _ := o.QueryTable("monitor").Count()
	return qs, nil
}

func UpdateMonitor(groupname string, hostname string,projectname string, scriptpath string, cron string, id int) Monitor {
	o := orm.NewOrm()
	var u Monitor
	err := o.Raw("update monitor set groupname= ?, projectname= ?, hostname = ?, scriptpath = ?, cron = ?  where id=?" , groupname,projectname,hostname,scriptpath,cron,id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}

func DelMonitor(id int) Monitor {
	o := orm.NewOrm()
	var u Monitor
	err := o.Raw("delete from monitor where id=?", id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}
func GetMonitorPag(pag int, count int) ([]Monitor, error) {
	o := orm.NewOrm()
	var u []Monitor
	_, err := o.Raw("select * from monitor limit ? offset ?",pag,count).QueryRows(&u)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func SelectMonitorHostname() []orm.Params {
	o := orm.NewOrm()

	var maps []orm.Params
	o.Raw("select  * from monitor").Values(&maps)

	//for _,term := range maps{
	//	fmt.Println(term["hostname"])
	//}
return maps
}

func SeachMonitor(key string, pag int, count int) ([]Monitor, error) {
	o := orm.NewOrm()
	var u []Monitor
	//_, err := o.Raw("select * from sshhost limit ? offset ?",pag,count).QueryRows(&u)
	_, err := o.Raw("select * from monitor where groupname like '%?%' limit ? offset ?",key,pag,count).QueryRows(&u)
	//select * from monitor where id like '%11%' limit 1,3
		if err != nil {
		return nil, err
	}

	return u, nil
}


//func SeachMonitorData(key string,pag int, count int) ([]Monitor, error) {
//	o := orm.NewOrm()
//	var u []Monitor
//	_, err := o.Raw("select * from monitor where groupname like '%?%' limit ? offset ?",key,pag,count).QueryRows(&u)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return u, nil
//}


func SeachMonitorCount( key string) int64 {
	orm := orm.NewOrm()
	qs := orm.QueryTable("monitor")
	n, _ := qs.Filter("groupname__contains", key).Count()
	return n
}

func SeachMonitorData(key string, pag int, count int) ([]Monitor, error) {
	orm := orm.NewOrm()
	var u []Monitor
	qs := orm.QueryTable("monitor")

	_, err:= qs.Filter("groupname__contains", key).Limit(pag,count).All(&u)

	if err != nil {
		return nil, err
	}
	return u, nil
}
