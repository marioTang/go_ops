package models

import "github.com/astaxie/beego/orm"

type Group struct {
	Id        int
	User string
	Passwd string
	Company string
	Merchants string
	Accesskeyid string
	Accesskeysecret string
	Remarks string
	Createtime string
}
func GetGroupUser(user string) Group {
	o := orm.NewOrm()
	var u Group
	err := o.Raw("select *  from group where user= ?", user).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}
func UploadUser(user string,passwd string, company string, merchants string, accesskeyid string, accesskeysecret string,remarks string,createtime string, id int) Group {
	o := orm.NewOrm()
	var u Group
	err := o.Raw("update group set user= ?,passwd = ?,company = ?,merchants = ? ,accesskeyid = ?,accesskeysecret = ?,remarks = ? ,createtime = ? where id=?" , user, passwd,company,merchants,accesskeyid,accesskeysecret,remarks,createtime,id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}