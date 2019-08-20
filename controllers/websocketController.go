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

type WebSocketController struct {
	beego.Controller
}

func (this *WebSocketController) Get() {
	tick := time.Tick(4 * time.Second)
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		beego.Error(err)
	}
	defer ws.Close()
	Clients[ws] = true
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
		hostinfo := <-Hostchan
		for client := range Clients {
			err := client.WriteJSON(hostinfo)
			if err != nil {
				beego.Error(err)
				client.Close()
				delete(Clients, client)
			}
		}
	}
}
