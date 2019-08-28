package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

var (
	SSHHosts = make(map[string]models.MachineSSH)
)

type MachineController struct {
	BaseController
}

func init() {
	SSHHosts = utils.SSHMapExport()
}

func (this *MachineController) Get() {
	this.TplName = "machine.html"
	this.Data["machine"] = SSHHosts
}

func (this *MachineController) Post() {
	this.TplName = "machine.html"
	saveMachine_btn := this.Input().Get("saveallmachines")
	exportMachine_btn := this.Input().Get("exportallmachines")
	if saveMachine_btn != "" {
		this.Save()
	} else if exportMachine_btn != "" {
		this.Save()
		this.Export()
	}
	this.Redirect("/machine", 301)

}

func (this *MachineController) Export() {
	message, err := utils.HostJsonGenerator()
	if err != nil {
		beego.Error(err)
	}
	beego.Info(message)
}

func (this *MachineController) Save() {
	nodenames := this.GetStrings("nodename")
	hostips := this.GetStrings("hostip")
	hostnames := this.GetStrings("hostname")
	users := this.GetStrings("username")
	passwords := this.GetStrings("password")
	keyfiles := this.GetStrings("keyfiles")
	sshports := this.GetStrings("sshport")
	authtypes := this.GetStrings("authtype")
	err := utils.HostSave(nodenames, hostips, hostnames, users, passwords, authtypes, sshports, keyfiles)
	if err != nil {
		beego.Error(err)
	}
	SSHHosts = utils.SSHMapExport()
}
