package utils

import (
	"net/smtp"
	"strings"
)

//func SendMail(addr string, a Auth, from string, to []string, msg []byte) error

/*!
username 发送者邮件
password 授权码
host 主机地址 smtp.qq.com:587 或 smtp.qq.com:25
to 接收邮箱 多个接收邮箱使用 ; 隔开
name 发送人名称
subject 发送主题
body 发送内容
mailType 发送邮件内容类型
*/
func SendMail(username, password, host, to, name,subject, body, mailType string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", username, password,hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + name + "<" + username + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, username, sendTo, msg)
	return err
}

//func main()  {
//	err := SendMail(
//		"zer******0115@foxmail.com",
//		"kcll******hjbbih",
//		"smtp.qq.com:587",
//		"**********@qq.com",
//		"学海无涯",
//		"测试",
//		"这是一封测试邮件",
//		"html",
//	)
//	if err != nil{
//		log.Fatal(err)
//	}
//}