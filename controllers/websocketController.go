package controllers

import (
	"net/http"
	"sync"
	"time"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var hostOnAndOff bool = true
var tick = time.Tick
var mux sync.Mutex

type HostWebSocketController struct {
	beego.Controller
}

type StepWebSocketController struct {
	beego.Controller
}

type ServiceWebSocketController struct {
	beego.Controller
}

type SSHWebSocketController struct {
	beego.Controller
}

func handleMessages(mux *sync.Mutex) {
	for {
		select {
		case hostinfo := <-Hostchan:
			for client := range HostClients {
				mux.Lock()
				err := client.WriteJSON(hostinfo)
				mux.Unlock()
				if err != nil {
					beego.Error(err)
					client.Close()
					delete(HostClients, client)
				}
			}
		case serviceinfo := <-Servicechan:
			for client := range ServiceClients {
				mux.Lock()
				err := client.WriteJSON(serviceinfo)
				mux.Unlock()
				if err != nil {
					beego.Error(err)
					client.Close()
					delete(ServiceClients, client)
				}
			}
		}
	}
}

func (this *HostWebSocketController) Get() {
	tick := time.Tick(5 * time.Second)
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	defer ws.Close()
	HostClients[ws] = true
	go handleMessages(&mux)
	for {
		if hostOnAndOff {
			hostinfo, err := utils.HostInfoRead()
			if err != nil {
				beego.Error(err)
				return
			}
			Hostchan <- hostinfo
			<-tick
		}
	}
}

func (this *ServiceWebSocketController) Get() {
	tick := time.Tick(20 * time.Second)
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	defer ws.Close()
	ServiceClients[ws] = true
	go handleMessages(&mux)
	for {
		for _, serviceEntity := range JsonStruct {
			serviceStatusUpdate, err := utils.ServiceInfo(serviceEntity)
			if err != nil {
				beego.Error(err)
				continue
			}
			Servicechan <- serviceStatusUpdate
		}
		<-tick
	}
}

func (this *SSHWebSocketController) Get() {
	this.TplName = "index.html"
	sshHost := SSHHosts[HostName]
	sshClient, err := utils.NewSshClient(sshHost)
	if err != nil {
		beego.Error(err)
	}
	defer sshClient.Close()
	// startTime := time.Now()
	sshConn, err := utils.NewSshConn(120, 32, sshClient)
	if err != nil {
		beego.Error(err)
	}
	defer sshConn.Close()
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	defer ws.Close()
	utils.SSHClients[ws] = true
	quitChan := make(chan bool, 3)
	go sshConn.ReceiveWsMsg(ws, quitChan)
	go sshConn.SendComboOutput(ws, quitChan)
	go sshConn.SessionWait(quitChan)
	<-quitChan
}

func (this *HostWebSocketController) StopHostSync() {
	hostOnAndOff = !hostOnAndOff
	beego.Info("I am trying to stop host sync")
	this.Redirect("/host", 302)
}
