package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Alibill struct {
	Id        int
	Users string
	Availableamount string
	Availablecashamount	 string
	Creditamount	 string
	Currency string
	Mybankcreditamount string
	Createtime time.Time

}

func InsertAlibill(alibill Alibill) error {
	o :=orm.NewOrm()
	u := Alibill{Availableamount: alibill.Availableamount,  Availablecashamount:alibill.Availablecashamount,Currency:alibill.Currency,Mybankcreditamount:alibill.Mybankcreditamount,
		Creditamount: alibill.Creditamount,	Createtime: alibill.Createtime,Users: alibill.Users,}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}
//
func GetAlibill(user string) Alibill {
	o := orm.NewOrm()
	var u Alibill
	err := o.Raw("select * from alibill where users= ?", user).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}
//
//
//func UpdateAsset( intranet_ip string, cpu_model string, hostname string, os string, cpu_num string, memory string, disk string) Asset {
//	o := orm.NewOrm()
//	var all2 Asset
//	err := o.Raw("update asset set cpu_model= ?,cpu_num = ?,os = ?,hostname = ? ,memory = ?,disk = ? where intranet_ip=?" , cpu_model,cpu_num,os,hostname,memory,disk,intranet_ip).QueryRow(&all2)
//	if err != nil {
//		return all2
//	}
//	return all2
//}
//
//
//func GetAssetCout() (int64, error) {
//	o := orm.NewOrm()
//	qs, _ := o.QueryTable("asset").Count()
//	return qs, nil
//}
//func GetAssetPag(pag int, count int) ([]Asset, error) {
//	o := orm.NewOrm()
//	var all1 []Asset
//	_, err := o.Raw("select * from asset limit ? offset ?",pag,count).QueryRows(&all1)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return all1, nil
//}
//
//func DelAsset(id int) Asset {
//	o := orm.NewOrm()
//	var all1 Asset
//	err := o.Raw("delete from asset where id=?", id).QueryRow(&all1)
//	if err != nil {
//		return all1
//	}
//	return all1
//}
//
//func SelectAssetinfo(id int) Asset {
//	o := orm.NewOrm()
//	var u Asset
//	err := o.Raw("select *  from asset where id= ?", id).QueryRow(&u)
//	if err != nil {
//		return u
//	}
//	return u
//}
