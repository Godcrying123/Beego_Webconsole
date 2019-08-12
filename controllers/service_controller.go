package controllers

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"path"
	"syscall"
	"time"
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

var jsonStruct map[string]models.Service

type ServiceController struct {
	beego.Controller
}

func (this *ServiceController) Get() {
	this.TplName = "service.html"
	//beego.Info(jsonStruct)
	this.Data["services"] = jsonStruct
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
	File, FileReader, _ := this.GetFile("importfile")
	ext := path.Ext(FileReader.Filename)
	// check the fileformat
	var AllowExtMap map[string]bool = map[string]bool{
		".json": true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		beego.Error("we do not support to import this type file")
		return
	}
	// Create Directory
	uploadDir := "static/upload/" + time.Now().Format("2006-01-02")
	oldMask := syscall.Umask(0)
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		beego.Error(err)
	}
	syscall.Umask(oldMask)
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext
	filePath := uploadDir + "/" + fileName
	defer File.Close()
	err = this.SaveToFile("importfile", filePath)
	if err != nil {
		beego.Error(err)
	}
	//beego.Info("Upload Successfully")
	jsonStruct, err = utils.ServicesJsonRead(filePath)
	//beego.Info(jsonStruct)
	this.Data["services"] = jsonStruct
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/service", 302)
}
