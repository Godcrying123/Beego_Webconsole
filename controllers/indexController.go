package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var (
	StepJsonStruct []models.MainSteps
	JsonStruct     map[string]models.Service
	ServiceClients = make(map[*websocket.Conn]bool)
	Servicechan    = make(chan models.Service)
	HostName       string
	SSHUrl         string
)

func init() {

}

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.TplName = "index.html"
	this.Data["stepsData"] = StepJsonStruct
	this.Data["services"] = JsonStruct
	HostName = this.Ctx.Input.Param(":hostname")
	if HostName != "" {
		SSHUrl = "/node/" + HostName + "/"
	} else {
		SSHUrl = "/"
	}
	// beego.Info(SSHUrl)
	if HostName == "" {
		HostName = "localhost"
	}
	this.Data["sshUrl"] = SSHUrl
	this.Data["machine"] = SSHHosts
	// beego.Info(HostName)
}

func (this *IndexController) Post() {
	this.TplName = "index.html"
	btn_stepimport := this.Input().Get("importallsteps")
	btn_serviceimport := this.Input().Get("importall")
	btn_hostonoff := this.Input().Get("syncoff")
	if btn_stepimport != "" {
		this.StepImport()
	} else if btn_serviceimport != "" {
		this.ServiceImport()
	} else if btn_hostonoff != "" {
		hostOnAndOff = !hostOnAndOff
	}
}

func (this *IndexController) StepImport() {
	filePath, err := this.FileUploadAndSave("importfilestep", ".json")
	if err != nil {
		beego.Error(err)
		return
	}
	StepJsonStruct, err = utils.StepJsonRead(filePath)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 302)
}

func (this *IndexController) ServiceImport() {
	filePath, err := this.FileUploadAndSave("importfile", ".json")
	if err != nil {
		beego.Error(err)
	}
	JsonStruct, err = utils.ServicesJsonRead(filePath)

	if err != nil {
		beego.Error(err)
	}
	for servicekey, service := range JsonStruct {
		serviceReturn, err := utils.ServiceInfo(service)
		if err != nil {
			beego.Error(err)
		}
		JsonStruct[servicekey] = serviceReturn

	}
	this.Redirect("/", 302)
}
