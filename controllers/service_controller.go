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

type ServiceController struct {
	beego.Controller
}

func (c *ServiceController) Get() {
	c.TplName = "service_upload.html"
}

func (this *ServiceController) Post1() {
	this.TplName = "service.html"

}

func (this *ServiceController) Post() {
	this.TplName = "service_upload.html"
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
	message, err := utils.Services_JsonGenerator(services)
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
	beego.Info("Upload Successfully")
	_, err = utils.Services_JsonRead(filePath)
	if err != nil {
		beego.Error(err)
	}
}
