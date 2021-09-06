package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type CheckLoginController struct {
	beego.Controller
	Read   bool
	Edit  bool
	Deldet bool
	Login bool
}

//
func (this *CheckLoginController) Prepare() {
	loginuser := this.GetSession("loginuser")
	if loginuser != nil {

	} else {this.Redirect("/login",302)}
	fmt.Println(loginuser)

}
