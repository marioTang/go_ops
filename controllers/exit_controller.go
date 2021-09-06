package controllers

type ExitController struct {
	CheckLoginController
}
func (this *ExitController)Get(){
	this.DelSession("loginuser")
	this.Redirect("/login",302)
}