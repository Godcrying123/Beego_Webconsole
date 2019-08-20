package controllers

import (
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

func (this *HostWebSocketController) Get() {
	tick := time.Tick(4 * time.Second)
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

func (this *ServiceWebSocketController) Get() {
	tick := time.Tick(100 * time.Second)
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
