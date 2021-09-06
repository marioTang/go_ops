package utils

import (
	"cmdb/models"
	"fmt"
	"github.com/astaxie/beego/toolbox"
)


//启动所有定时任务
func Allcron() {

	hostnames := models.SelectMonitorHostname()

	for _, term := range hostnames {
		crontab := term["cron"].(string)
		taskname := term["projectname"].(string)
		host := term["hostname"].(string)
		fmt.Println(host)
		cmdexec := term["scriptpath"].(string)
		publcikey := models.SelctSshhostId("119.3.106.91").Publcikey

		tk := toolbox.NewTask(taskname, crontab, func() error {

			fmt.Println(string(Runshell(host, "root", "", 22, publcikey, cmdexec)));
			return nil
		})
		err := tk.Run()
		if err != nil {
			fmt.Println(err)
		}

		toolbox.AddTask(taskname, tk)
		toolbox.StartTask()
	}

}

//添加定时任务
func Addcron(host string) {

	hostname := models.SelctMonitorId(host).Hostname
	taskname := models.SelctMonitorId(host).Projectname
	cmdexec := models.SelctMonitorId(host).Scriptpath
	publcikey := models.SelctSshhostId("119.3.106.91").Publcikey
	crontab   := models.SelctMonitorId(host).Cron

	tk := toolbox.NewTask(taskname, crontab, func() error {
		fmt.Println(string(Runshell(hostname, "root", "", 22, publcikey, cmdexec)));
		return nil
	})
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}

	toolbox.AddTask(taskname, tk)
	toolbox.StartTask()
}


//删除定时任务
func Delcron(taskname string) {

	toolbox.DeleteTask(taskname)
}


