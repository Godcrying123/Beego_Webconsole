package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type ServiceController struct {
	BaseController
}

func (this *ServiceController) Get() {
	this.TplName = "service.html"
	if len(JsonStruct) != 0 {
		this.Data["serviceExist"] = true
		this.Data["services"] = JsonStruct
	} else {
		this.Data["serviceExist"] = false
	}
	this.Data["stepList"] = StepJsonStruct
	this.Data["sshUrl"] = SSHUrl
}

func (this *ServiceController) Post() {
	this.TplName = "service.html"
	this.Export()
}

func (this *ServiceController) Export() {
	var services = make(map[string]models.Service)
	servicesname := this.GetStrings("service_name")
	servicesversions := this.GetStrings("service_version")
	for index := 0; index < len(servicesname); index++ {
		var servicestruct = models.Service{
			ServiceName:    servicesname[index],
			ServiceVersion: servicesversions[index],
		}
		services[servicesname[index]] = servicestruct
	}
	message, err := utils.ServicesJsonGenerator(services)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(message)
}
