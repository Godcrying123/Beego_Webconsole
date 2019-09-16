package controllers

import (
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type StepController struct {
	BaseController
}

func (this *StepController) Get() {
	this.TplName = "step.html"
	if len(StepJsonStruct) != 0 {
		this.Data["stepExist"] = true
		this.Data["stepList"] = StepJsonStruct
	} else {
		this.Data["stepExist"] = false
	}
	this.Data["services"] = JsonStruct
	this.Data["sshUrl"] = SSHUrl
	beego.Info(SSHUrl)
}

func (this *StepController) Post() {
	this.TplName = "step.html"
	this.Export()
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
