package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var (
	jsonStruct     = make(map[string]models.Service)
	ServiceClients = make(map[*websocket.Conn]bool)
	Servicechan    = make(chan models.Service)
)

type ServiceController struct {
	BaseController
}

func (this *ServiceController) Get() {
	this.TplName = "service_upload.html"
	this.Data["services"] = jsonStruct
}

func (this *ServiceController) Upload() {
	this.TplName = "service.html"
	if len(jsonStruct) != 0 {
		this.Data["serviceExist"] = true
		this.Data["services"] = jsonStruct
	} else {
		this.Data["serviceExist"] = false
	}
}

func (this *ServiceController) Post() {
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
	}
	jsonStruct, err = utils.ServicesJsonRead(filePath)

	if err != nil {
		beego.Error(err)
	}
	for servicekey, service := range jsonStruct {
		serviceReturn, err := utils.ServiceInfo(service)
		if err != nil {
			beego.Error(err)
		}
		jsonStruct[servicekey] = serviceReturn

	}
	this.Redirect("/service", 302)
}
