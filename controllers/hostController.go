package controllers

import (
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type HostController struct {
	beego.Controller
}

func (this *HostController) Get() {
	this.TplName = "host.html"
	hostinfo, err := utils.HostInfoRead()
	if err != nil {
		this.Redirect("/host", 302)
		return
	}
	this.Data["hostinfo"] = hostinfo
}
