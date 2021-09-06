package routers

import (
	"cmdb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	//验证码
	beego.Router("/captcha", &controllers.CaptchaController{})
	beego.Router("/test", &controllers.TestController{}, "get:Test")
	beego.Router("/test", &controllers.TestController{}, "post:Test")

	beego.Router("/test",&controllers.TestController{},"get:Test")
	beego.Router("/test",&controllers.TestController{},"post:Test")
	beego.Router("/test",&controllers.TestController{},"put:Test")
	beego.Router("/test",&controllers.TestController{},"delete:Test")



	//beego.Router("/exit", &controllers.ExitController{})
	//beego.Router("/assetinfo", &controllers.AssetinfoController{})
	//beego.Router("/api", &controllers.ApiController{})
	//beego.Router("/", &controllers.IndexController{})
	//beego.Router("/aliasset", &controllers.AliController{})
	beego.Router("/ssh", &controllers.SshController{})
	//beego.Router("/sshinfo", &controllers.HostinfoController{})
	//beego.Router("/monitorinfo", &controllers.SeachController{})
	//beego.Router("/updatamonitor", &controllers.Updatamonitor{})
	//
	//beego.Router("/sshrun", &controllers.RunSshController{})
	//beego.Router("/file", &controllers.FileController{})
	//beego.Router("/group", &controllers.GroupController{})
	//beego.Router("/plmxs", &controllers.PlxmxsController{})
	beego.Router("/plmxsversion", &controllers.PlxmxsVersionController{})



	//beego.Router("/je", &controllers.AlibillController{})
	//beego.Router("/adduser", &controllers.AdduserController{})
	//beego.Router("/monitor", &controllers.MonitorController{})
	//beego.Router("/seach", &controllers.SeachController{})
	//beego.Router("/userinfo", &controllers.UserinfoController{})
	//beego.Router("/fileupload", &controllers.FileUploadController{})
	//beego.Router("/ws/ssh", &controllers.WebSocketController{})
	//beego.Router("/wsssh", &controllers.SshlistController{})






	//beego.Router("/websocket", &controllers.IndexController{})
	//beego.Router("/ws", &controllers.MyWebSocketController{})


}
