package controllers

import (
	"bytes"
	"cmdb/logger"
	"cmdb/models"
	"cmdb/ssh"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)


type WebSocketController struct {
	beego.Controller
}
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024 * 1024 * 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handle webSocket connection.
// first,we establish a ssh connection to ssh server when a webSocket comes;
// then we deliver ssh data via ssh connection between browser and ssh server.
// That is, read webSocket data from browser (e.g. 'ls' command) and send data to ssh server via ssh connection;
// the other hand, read returned ssh data from ssh server and write back to browser via webSocket API.
func (c *WebSocketController) Get()  {
	cols, _ := strconv.Atoi(c.Input().Get("cols"))
	rows, _ := strconv.Atoi(c.Input().Get("rows"))
	id, _ := strconv.Atoi(c.Input().Get("hostid"))

	fmt.Println(cols)
	fmt.Println(id)
	fmt.Println(rows)

	wsConn, err := upGrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		logger.Logger.Println(err)
	}

	defer wsConn.Close()
	host := models.Getsshws_data_id(id)
	client, err := ssh.NewClient(host.Addr, host.Port, host.Username, "", host.Publcikey)
	if err != nil {
		logger.Logger.Println(err)
	}


	ssConn, err := ssh.NewSshConn(cols, rows, client.SSHClient)

	defer ssConn.Close()

	quitChan := make(chan bool, 3)

	var logBuff = new(bytes.Buffer)
	go ssConn.ReceiveWsMsg(wsConn, logBuff, quitChan)
	go ssConn.SendComboOutput(wsConn, quitChan)
	go ssConn.SessionWait(quitChan)

	<-quitChan
	logrus.Info("websocket finished")
}