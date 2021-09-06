package models

import (
	"github.com/astaxie/beego/orm"
)

type Ecsinfo struct {
	Id int
	Instanceid string
	Instancename string
	Instancetype string
	Status string
	Regionid string
	Creationtime  string
	Hostname string
	Corecount int
	Threadspecore  int
	Internetchargetype  string
	Bandwidth  int
	Ipaddress  string
	Inneripaddress  string
	Publicipaddress string
	Instancenetworktype string
	Localstorageamount int
	Localstoragecapacity string
	Memory int
	Osname string
	Ostype string

}
type Count struct {
	Countdata  int64
}
func InsertEscinfo(escinfo Ecsinfo) error {
	o :=orm.NewOrm()
	u := Ecsinfo{Instanceid: escinfo.Instanceid, Instancename:escinfo.Instancename, Instancetype: escinfo.Instancetype, Status: escinfo.Status, Regionid: escinfo.Regionid, Creationtime: escinfo.Creationtime,Hostname: escinfo.Hostname, Corecount:escinfo.Corecount, Threadspecore :escinfo.Threadspecore,Internetchargetype: escinfo.Internetchargetype ,Bandwidth:escinfo.Bandwidth,Ipaddress:escinfo.Ipaddress,Inneripaddress :escinfo.Inneripaddress,Publicipaddress:escinfo.Publicipaddress,Instancenetworktype:escinfo.Instancenetworktype,Localstorageamount:escinfo.Localstorageamount ,Localstoragecapacity: escinfo.Localstoragecapacity,Memory:escinfo.Memory,Osname:escinfo.Osname,Ostype:escinfo.Ostype}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func SelectInstanceid(Id string) Ecsinfo {
	o := orm.NewOrm()
	var u Ecsinfo
	err := o.Raw("select instanceid  from ecsinfo where instanceid= ?", Id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}


//func GetAssetCout() (int64, error) {
//	o := orm.NewOrm()
//	qs, _ := o.QueryTable("asset").Count()
//	return qs, nil
//}

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
func SelectAliAssetinfo(Regionid string) Ecsinfo {
	o := orm.NewOrm()
	var u Ecsinfo
	err := o.Raw("select *  from ecsinfo where regionid= ?", Regionid).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}
func GetAliAssetPag(regionid string, pag int, count int ) ([]Ecsinfo, error) {
	o := orm.NewOrm()
	var all1 []Ecsinfo
	_, err := o.Raw("select * from ecsinfo where ecsinfo.regionid = ? limit ? offset ?",regionid, pag, count).QueryRows(&all1)

	if err != nil {
		return nil, err
	}

	return all1, nil
}
//func GetAliAssetCount(regionid string) int64 {
//	o := orm.NewOrm()
//	//var all1 []Count
//	cnt, err := o.QueryTable("ecsinfo").Filter(regionid).Count()
//	//_, err := o.Raw("select count(*) from ecsinfo where ecsinfo.regionid = ?", regionid).QueryRows(&all1)
//	if err !=nil {
//
//	}
//	return cnt
//}
func GetAliAssetCount(regionid string) (count int64) {
	o := orm.NewOrm()
	qs :=o.QueryTable(new(Ecsinfo))
	count,_=qs.Filter("regionid",regionid).Count()
	return count
}