package controllers

import (
	"webconsole_sma/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var (
	turnOnOff   bool
	onAndOff    = make(chan bool, 1)
	HostClients = make(map[*websocket.Conn]bool)
	Hostchan    = make(chan models.Machine)
)

type HostController struct {
	beego.Controller
}

func (this *HostController) Get() {
	this.TplName = "host.html"
	// select {
	// case turnOnOff := <-onAndOff:
	// 	this.Data["switch"] = turnOnOff
	// default:
	// 	this.Data["switch"] = false
	// }
	// hostinfo, err := utils.HostInfoRead()
	// if err != nil {
	// 	this.Redirect("/host", 302)
	// 	return
	// }
	// this.Data["hostinfo"] = hostinfo
}

func (this *HostController) Post1() {
	// this.TplName = "host.html"
	// tick := time.Tick(2 * time.Second)
	// syncoff := this.Input().Get("syncoff")
	// syncon := this.Input().Get("syncon")
	// if syncoff != "" {
	// 	onAndOff <- false
	// 	// ws.Close()
	// 	this.Redirect("/host", 302)
	// }
	// if syncon != "" {
	// 	onAndOff <- true
	// 	// clients[ws] = true
	// 	// defer ws.Close()
	// 	go func() {
	// 		for {
	// 			hostinfo := <-hostchan
	// 			for client := range clients {
	// 				err := client.WriteJSON(hostinfo)
	// 				if err != nil {
	// 					log.Printf("client.WriteJSON error: %v", err)
	// 					client.Close()
	// 					delete(clients, client)
	// 				}
	// 			}

	// 		}
	// 	}()
	// 	for {
	// 		hostinfo, err := utils.HostInfoRead()
	// 		if err != nil {
	// 			this.Redirect("/host", 302)
	// 			return
	// 		}
	// 		beego.Info("I am waiting")
	// 		hostchan <- hostinfo
	// 		<-tick
	// 	}
	// }
}
