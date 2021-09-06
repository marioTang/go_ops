package models

import (
	"github.com/astaxie/beego/orm"
)

type Regions struct {
	Id        int  `json:"id" gorm:"column:id"`
	Regionid string `json:"regionId" gorm:"column:regionId"`
	Regionendpoint string `json:"regionEndpoint"`
	Localname string `json:"localname"`

}
func InsertRegions(regions Regions) error {
	o :=orm.NewOrm()
	u := Regions{Regionid: regions.Regionid, Regionendpoint:regions.Regionid, Localname:regions.Localname}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}




func SelectRegionId(Id string) Regions {
	o := orm.NewOrm()
	var u Regions
	err := o.Raw("select regionid  from regions where regionid= ?", Id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}

func SelectRegionsData() string {
	o := orm.NewOrm()
	var u string
	_, err := o.Raw("select regionid  from regions").QueryRows(&u)
	if err != nil {
     return u
	}
	return u
}


