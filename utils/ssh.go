package utils

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

//效验公钥
func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	//keyPath, err := homedir.Expand(kPath)
	//if err != nil {
	//	log.Fatal("find key's home dir failed", err)
	//}
	//key, err := ioutil.ReadFile(keyPath)
	//if err != nil {
	//	log.Fatal("ssh key file read failed", err)
	//}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey([]byte(kPath))
	if err != nil {
		log.Println("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

//执行命令
func Runshell(sshhost, username, password string, sshport int, publickey, ssharg string) []byte {

	config := &ssh.ClientConfig{
		Timeout:         time.Second,//ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	if len(password) == 0 {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(publickey)}

	} else {
		config.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshhost, sshport)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("创建ssh client 失败",err)
	}
	defer sshClient.Close()


	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Println("创建ssh session 失败",err)
	}
	defer session.Close()
	//执行远程命令
	combo,err := session.CombinedOutput(ssharg)
	if err != nil {
		return nil

	}
	return combo

	log.Println("命令输出:",string(combo))
	return nil
}

//上传文件
func Uploadfile(sshhost ,username,password string, sshport int,publickey string,localFilePath string, remotePath string){

	config := &ssh.ClientConfig{
		Timeout:         time.Second,//ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	if len(password) == 0 {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(publickey)}

	} else {
		config.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshhost, sshport)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("创建ssh client 失败",err)
	}
	defer sshClient.Close()
	sftpClient,_:= sftp.NewClient(sshClient)


	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Println("创建ssh session 失败",err)
	}
	defer session.Close()

	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		fmt.Println(err)

	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)

	dstFile, err := sftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join(remotePath, remoteFileName))
		fmt.Println(err)

	}
	defer dstFile.Close()

	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localFilePath)
		fmt.Println(err)

	}
	dstFile.Write(ff)
	fmt.Println(localFilePath + "  copy file to remote server finished!")


}

func Download(sshhost ,username,password string, sshport int,publickey string,dspath string, srcpath string)(srcPath, dstPath string) {

	config := &ssh.ClientConfig{
		Timeout:         time.Second,//ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	if len(password) == 0 {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(publickey)}

	} else {
		config.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshhost, sshport)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("创建ssh client 失败",err)
	}
	defer sshClient.Close()
	sftpClient,_:= sftp.NewClient(sshClient)

	srcFile, _ := sftpClient.Open(srcpath) //远程
	dstFile, _ := os.Create(dspath) //本地
	defer func(){
		_ = srcFile.Close()
		_ = dstFile.Close()
	}()

	if _, err := srcFile.WriteTo(dstFile); err != nil {
		fmt.Println("error occurred", err)
	}
	fmt.Println("文件下载完毕")
	return
}

func Ifssh(sshhost, username, password string, sshport int, publickey string) (error) {

	config := &ssh.ClientConfig{
		Timeout:         time.Second,//ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}

	if len(password) == 0 {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(publickey)}
	} else {
		config.Auth = []ssh.AuthMethod{ssh.Password(password)}
	}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshhost, sshport)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		fmt.Println("创建ssh client 失败",err)
	}
	defer sshClient.Close()


	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Println("创建ssh session 失败",err)
		return error(err)
	}
	defer session.Close()
	return nil
}