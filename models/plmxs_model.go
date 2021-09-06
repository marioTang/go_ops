package models

import (
	"github.com/astaxie/beego/orm"
)

type Upload struct {
	Id        int
	Host  string
	Filename  string
	Sourcefile string
	Filepath  string
	Source      string
	Createtime string
}

func PlmxsUpload(upload Upload ) error {
	o :=orm.NewOrm()
	u := Upload{0,upload.Host,upload.Filename,upload.Sourcefile,upload.Filepath,upload.Source,upload.Createtime}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}
func CoutPlmxs() (int64, error) {
	o := orm.NewOrm()
	qs, _ := o.QueryTable("upload").Count()
	return qs, nil
}
//
//func SelctSshhostId(hostname string) Sshhost {
//	var u Sshhost
//	o :=orm.NewOrm()
//	err := o.Raw("select * from sshhost where hostname = ?",hostname).QueryRow(&u)
//	if err != nil {
//	}
//	return u
//}
//
func SelectCoutPlmxs() ([]Upload) {
	o := orm.NewOrm()
	var u []Upload
	o.QueryTable("upload").All(&u)
	return u
}
//
//func CoutSshhost() (int64, error) {
//	o := orm.NewOrm()
//	qs, _ := o.QueryTable("sshhost").Count()
//	return qs, nil
//}
func SelctUpload(id int) Upload {
	var u Upload
	o :=orm.NewOrm()
	err := o.Raw("select * from upload where id = ?",id).QueryRow(&u)
	if err != nil {
	}
	return u
}
func SelctUploadHost(host string) Upload {
	var u Upload
	o :=orm.NewOrm()
	err := o.Raw("select * from upload where host = ?",host).QueryRow(&u)
	if err != nil {
	}
	return u
}

func DelUpload(id int) Upload {
	o := orm.NewOrm()
	var u Upload
	err := o.Raw("delete from upload where id=?", id).QueryRow(&u)
	if err != nil {
		return u
	}
	return u
}
func GetUploadPag(pag int, count int) ([]Upload, error) {
	o := orm.NewOrm()
	var u []Upload
	_, err := o.Raw("select * from upload limit ? offset ?",pag,count).QueryRows(&u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
