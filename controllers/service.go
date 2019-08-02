package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type ServiceController struct {
	beego.Controller
}

func (c *ServiceController) Get() {
	c.TplName = "service.html"
}

func (this *ServiceController) Post1() {
	this.TplName = "service.html"

}

func (this *ServiceController) Post() {
	this.TplName = "service.html"
	btn_import := this.Input().Get("importall")
	btn_export := this.Input().Get("exportall")
	//beego.Info(btn_export)
	//beego.Info(btn_import)
	if btn_export != "" {
		this.Export()
	} else if btn_import != "" {
		this.Import()
		beego.Info(btn_import)
	}
}

func (this *ServiceController) Export() {
	var services = make(map[string]models.Service)
	services_name := this.GetStrings("service_name")
	services_versions := this.GetStrings("service_version")
	for index := 0; index < len(services_name); index++ {
		var service_struct = models.Service{
			ServiceName:    services_name[index],
			ServiceVersion: services_versions[index],
		}
		services[services_name[index]] = service_struct
	}
	message, err := utils.Services_XMLGenerator(services)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(message)
}

func (this *ServiceController) Import() {

}
