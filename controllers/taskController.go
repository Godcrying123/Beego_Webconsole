package controllers

import (
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

type TaskController struct {
	BaseController
}

func init() {
}

func (this *TaskController) Get() {
	// this.TplName = "autotask.html"
	this.TplName = "task.html"
	if len(TaskJsonMap) != 0 {
		this.Data["taskExist"] = true
		this.Data["taskData"] = TaskJsonMap
	} else {
		this.Data["taskExist"] = false
	}
	this.Data["stepList"] = StepJsonStruct
	this.Data["services"] = JsonStruct
	this.Data["machine"] = SSHHosts
	this.Data["sshUrl"] = SSHUrl
}

func (this *TaskController) Post() {
	// this.TplName = "autotask.html"
	this.TplName = "task.html"
	// btn_exportTask := this.Input().Get("exportalltask")
	this.Export()
	// this.Redirect("/task", 302)
}

func (this *TaskController) Export() {
	taskname := this.GetStrings("task_name")
	tasksummary := this.GetStrings("task_summary")
	tasknode := this.GetStrings("task_nodes")
	taskcommand := this.GetStrings("task_commands")
	maintaskslices := utils.TaskAnalyzer(taskname, tasksummary, tasknode, taskcommand)
	_, err := utils.TaskJsonGenerator(maintaskslices)
	if err != nil {
		beego.Error(err)
	}
}
