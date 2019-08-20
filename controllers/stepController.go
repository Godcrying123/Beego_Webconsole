package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

var stepJsonStruct []models.MainSteps

type StepController struct {
	BaseController
}

func (this *StepController) Get() {
	this.TplName = "step_upload.html"
	this.Data["stepsData"] = stepJsonStruct
}

func (this *StepController) Post() {
	this.TplName = "step_upload.html"
	beego.Info("Running")
	this.Import()
}

func (this *StepController) Edit() {
	this.TplName = "step.html"
	if len(stepJsonStruct) != 0 {
		this.Data["stepExist"] = true
		this.Data["stepList"] = stepJsonStruct
		// beego.Info(stepJsonStruct)
	} else {
		this.Data["stepExist"] = false
	}
}

func (this *StepController) Import() {
	filePath, err := this.FileUploadAndSave("importfilestep", ".json")
	if err != nil {
		beego.Error(err)
		return
	}
	stepJsonStruct, err = utils.StepJsonRead(filePath)
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
	mainstepslice := utils.StepAnalyzer(mainstep, stepname, stepsummary, stepcommand)
	_, err := utils.StepJsonGenerator(mainstepslice)
	if err != nil {
		beego.Error(err)
	}
}
