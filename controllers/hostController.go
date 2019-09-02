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
	this.TplName = "host_test.html"
	beego.Info(hostOnAndOff)
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

func (this *HostController) Post() {
	this.TplName = "host_test.html"
	hostOnAndOff = !hostOnAndOff
	// this.Redirect("/host", 302)
}
