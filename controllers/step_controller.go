package controllers

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"path"
	"syscall"
	"time"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type StepController struct {
	beego.Controller
}

func (this *StepController) Get() {
	this.TplName = "step_upload.html"
}

func (this *StepController) Post() {
	this.TplName = "step_upload.html"
	this.Import()
}

func (this *StepController) Import() {
	File, FileReader, _ := this.GetFile("importfilestep")
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
	err = this.SaveToFile("importfilestep", filePath)
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
	this.Redirect("/step", 302)
}

func (this *StepController) Export() {
	mainstep := this.GetStrings("main_step")
	stepname := this.GetStrings("step_name")
	stepsummary := this.GetStrings("step_summary")
	stepcommand := this.GetStrings("step_command")
	mainstepsmap := utils.StepAnalyzer(mainstep, stepname, stepsummary, stepcommand)
	_, err := utils.StepJsonGenerator(mainstepsmap)
	if err != nil {
		beego.Error(err)
	}
}
