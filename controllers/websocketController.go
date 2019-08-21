package controllers

import (
	"bytes"
	"net/http"
	"time"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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

func (this *HostWebSocketController) Get() {
	tick := time.Tick(4 * time.Second)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	defer ws.Close()
	HostClients[ws] = true
	go handleMessages()
	for {
		hostinfo, err := utils.HostInfoRead()
		if err != nil {
			beego.Error(err)
			return
		}
		Hostchan <- hostinfo
		<-tick
	}
}

func (this *ServiceWebSocketController) Get() {
	tick := time.Tick(20 * time.Second)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	defer ws.Close()
	ServiceClients[ws] = true
	go handleMessages()
	for {
		for _, serviceEntity := range jsonStruct {
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
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	defer ws.Close()
	utils.SSHClients[ws] = true
	go handleMessages()
	var byteBuilder bytes.Buffer
	for {
		for client := range utils.SSHClients {
			_, input, err := client.ReadMessage()
			if err != nil {
				beego.Error(err)
				client.Close()
				delete(utils.SSHClients, client)
			}
			if bytes.Equal(input, []byte{13}) {
				beego.Info(byteBuilder.String())
				byteBuilder.Reset()
			} else {
				byteBuilder.Write(input)
			}
		}
	}
}

func handleMessages() {
	for {
		select {
		case hostinfo := <-Hostchan:
			for client := range HostClients {
				err := client.WriteJSON(hostinfo)
				if err != nil {
					beego.Error(err)
					client.Close()
					delete(HostClients, client)
				}
			}
		case serviceinfo := <-Servicechan:
			for client := range ServiceClients {
				err := client.WriteJSON(serviceinfo)
				if err != nil {
					beego.Error(err)
					client.Close()
					delete(ServiceClients, client)
				}
			}
		}
	}
}
