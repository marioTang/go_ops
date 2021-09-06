package models

import (
	"github.com/astaxie/beego/orm"
)

type Sshhost struct {
	Id        int
	Hostname  string
	Username  string
	Password  string
	Port      int
	Publcikey string
}

func InsertSshhost(sshhost Sshhost ) error {
	o :=orm.NewOrm()
	u := Sshhost{0,sshhost.Hostname,sshhost.Username,sshhost.Password,sshhost.Port,sshhost.Publcikey}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

//
func SelctSshhostId(hostname string) Sshhost {
	var u Sshhost
	o :=orm.NewOrm()
	err := o.Raw("select * from sshhost where hostname = ?",hostname).QueryRow(&u)
	if err != nil {
	}
	return u
}

func SelectSshhost() ([]Sshhost) {
	o := orm.NewOrm()
	var u []Sshhost
	o.QueryTable("sshhost").All(&u)
	return u
}

func CoutSshhost() (int64, error) {
	o := orm.NewOrm()
	qs, _ := o.QueryTable("sshhost").Count()
	return qs, nil
}
func SelctSshhost(id int) Sshhost {
	var u Sshhost
	o :=orm.NewOrm()
	err := o.Raw("select * from sshhost where id = ?",id).QueryRow(&u)
	if err != nil {
	}
	return u
}
func UpdateSshhost(hostname string,username string, password string, port int, publcikey string, id int) Sshhost {
	o := orm.NewOrm()
	var u Sshhost
	err := o.Raw("update sshhost set hostname= ?, username = ?, password = ?,port = ?, publcikey = ?  where id=?" , hostname,username,password,port,publcikey,id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}

func DelSshhost(id int) Sshhost {
	o := orm.NewOrm()
	var u Sshhost
	err := o.Raw("delete from sshhost where id=?", id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}
func GetSshhostPag(pag int, count int) ([]Sshhost, error) {
	o := orm.NewOrm()
	var u []Sshhost
	_, err := o.Raw("select * from sshhost limit ? offset ?",pag,count).QueryRows(&u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
