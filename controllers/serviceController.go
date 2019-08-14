package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

var jsonStruct map[string]models.Service

type ServiceController struct {
	BaseController
}

func (this *ServiceController) Get() {
	this.TplName = "service.html"
	//beego.Info(jsonStruct)
	this.Data["services"] = jsonStruct
}

func (this *ServiceController) Upload() {
	this.TplName = "service_upload.html"
}

func (this *ServiceController) Post() {
	this.TplName = "service.html"
	btn_import := this.Input().Get("importall")
	btn_export := this.Input().Get("exportall")
	if btn_export != "" {
		this.Export()
	} else if btn_import != "" {
		this.Import()
	}
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

func (this *ServiceController) Import() {
	filePath, err := this.FileUploadAndSave("importfile", ".json")
	if err != nil {
		beego.Error(err)
		return
	}
	jsonStruct, err = utils.ServicesJsonRead(filePath)
	this.Data["services"] = jsonStruct
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/service", 302)
}
