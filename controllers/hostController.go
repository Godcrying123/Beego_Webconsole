package controllers

import (
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type HostController struct {
	beego.Controller
}

func (this *HostController) Get() {
	this.Data["switch"] = false
	this.TplName = "host.html"
	hostinfo, err := utils.HostInfoRead()
	if err != nil {
		this.Redirect("/host", 302)
		return
	}
	this.Data["hostinfo"] = hostinfo
}

func (this *HostController) Post() {
	this.TplName = "host.html"
	this.Data["switch"] = true
	// tick := time.Tick(2 * time.Second)
	// var hostchan = make(chan models.Machine)
	var onAndoff = make(chan bool)
	syncoff := this.Input().Get("syncoff")
	syncon := this.Input().Get("syncon")
	beego.Info(syncoff)
	beego.Info(syncon)
	if syncon != "" {
		beego.Info("Sync On")
		onAndoff <- true
	}
	beego.Info("I am going here")
	status := <-onAndoff
	beego.Info(status)
	// go func() {
	// 	for {
	// 		hostinfo := <-hostchan
	// 		// this.Data["hostinfo"] = hostinfo
	// 		beego.Info(hostinfo.Memory.UsedMemory)
	// 	}
	// }()
	// for {
	// 	hostinfo, err := utils.HostInfoRead()
	// 	if err != nil {
	// 		this.Redirect("/host", 302)
	// 		return
	// 	}
	// 	beego.Info("I am waiting")
	// 	hostchan <- hostinfo
	// 	<-tick
	// }
	this.Redirect("/host", 302)

}
