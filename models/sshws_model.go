package models

import (
	"github.com/astaxie/beego/orm"
)

type Sshws struct {
	Id        int
	Addr  string
	Username  string
	Password  string
	Port      int
	Publcikey string
}

//func InsertSshhost(sshhost Sshhost ) error {
//	o :=orm.NewOrm()
//	u := Sshhost{0,sshhost.Hostname,sshhost.Username,sshhost.Password,sshhost.Port,sshhost.Publcikey}
//	_, err := o.Insert(&u)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//
func Getsshws_data_addr(addr string) Sshws {
	var u Sshws
	o :=orm.NewOrm()
	err := o.Raw("select * from sshws where addr = ?",addr).QueryRow(&u)
	if err != nil {
	}
	return u
}

func Getsshws_array_all() ([]Sshws) {
	o := orm.NewOrm()
	var u []Sshws
	o.QueryTable("sshws").All(&u)
	return u
}

func Coutsshws() (int64, error) {
	o := orm.NewOrm()
	qs, _ := o.QueryTable("sshws").Count()
	return qs, nil
}
func Getsshws_data_id(id int) Sshws {
	var u Sshws
	o :=orm.NewOrm()
	err := o.Raw("select * from sshws where id = ?",id).QueryRow(&u)
	if err != nil {
	}
	return u
}
