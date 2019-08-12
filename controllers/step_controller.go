package controllers

import (
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type StepController struct {
	beego.Controller
}

func (this *StepController) Get() {
	this.TplName = "step.html"
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
	mainstepsmap := utils.StepAnalyzer(mainstep, stepname, stepsummary, stepcommand)
	_, err := utils.StepJsonGenerator(mainstepsmap)
	if err != nil {
		beego.Error(err)
	}
}
