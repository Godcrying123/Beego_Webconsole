package controllers

import (
	"strconv"
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type HostController struct {
	beego.Controller
}

func (this *HostController) Get() {
	this.TplName = "host.html"
	var host models.Machine
	hostinfo, err := utils.HostInfoRead(host)
	if err != nil {
		this.Redirect("/host", 302)
		return
	}
	this.Data["CPUutilizations"] = hostinfo.CPU.CPUPercentage
	this.Data["hostinfocpu"] = []string{hostinfo.CPU.CPUModelandFrequency, strconv.Itoa(hostinfo.CPU.CPUCores)}
	this.Data["hostinfoMemory"] = hostinfo.Memory
	this.Data["hostinfoDisk"] = hostinfo.DiskSpace
	this.Data["hostname"] = hostinfo.HostName
	this.Data["os"] = hostinfo.OS
}
