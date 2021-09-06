package models

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id int
	Name           string
	Passwd         string
	Projectteam    string

}



func CoutUser() (int64, error) {
	o := orm.NewOrm()
	qs, _ := o.QueryTable("users").Count()
	return qs, nil
}

func GetUserPag(pag int, count int) ([]Users, error) {
	o := orm.NewOrm()
	var all1 []Users
	_, err := o.Raw("select * from users limit ? offset ?",pag,count).QueryRows(&all1)

	if err != nil {
		return nil, err
	}

	return all1, nil
}

func GetUserID(username string, password string) (int) {
	var u Users
	o :=orm.NewOrm()
	err := o.Raw("select id from users where name = ? and passwd = ?",username, password).QueryRow(&u)
	if err != nil {
	}
	return u.Id
}

func InsertUser(users Users) error {
	o :=orm.NewOrm()
	u := Users{
		Name: users.Name,
		Passwd:users.Passwd,
		Projectteam:users.Projectteam,
	}
	_, err := o.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}
func GetUserName(name string) Users {
	var u Users
	o :=orm.NewOrm()
	err := o.Raw("select name from users where name = ?",name).QueryRow(&u)
	if err != nil {
	}
	return u
}

func GetUserData(id int) Users {
	var u Users
	o :=orm.NewOrm()
	err := o.Raw("select * from users where id = ?",id).QueryRow(&u)
	if err != nil {
	}
	return u
}
func DelUser(id int) Users {
	var u Users
	o :=orm.NewOrm()
	err := o.Raw("select * from users where id = ?",id).QueryRow(&u)
	if err != nil {
	}
	return u
}