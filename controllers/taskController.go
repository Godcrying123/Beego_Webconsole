package controllers

import (
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

var (
	TaskJsonMap map[string]models.MainTasks
)

type TaskController struct {
	BaseController
}

func (this *TaskController) Get() {
	// this.TplName = "autotask.html"
	this.TplName = "task_run.html"
	this.Data["taskData"] = TaskJsonMap
}

func (this *TaskController) Post() {
	this.TplName = "task_run.html"
	btn_exportTask := this.Input().Get("exportalltask")
	btn_importTask := this.Input().Get("importalltasks")
	btn_runTask := this.Input().Get("runTask")
	btn_runAllTask := this.Input().Get("runAllTask")
	btn_taskDetail := this.Input().Get("AllTaskDetails")
	task_node := this.Input().Get("TaskNode")
	task_command := this.Input().Get("TaskCommand")
	beego.Info(btn_taskDetail)
	beego.Info(task_node)
	beego.Info(btn_runTask)
	beego.Info(btn_runAllTask)
	beego.Info(task_command)
	if btn_exportTask != "" {
		this.Export()
	} else if btn_importTask != "" {
		beego.Info(btn_importTask)
		this.Import()
	} else if btn_runTask != "" {
		eachTask := models.EachTask{}
		if task_node != "" && task_command != "" {
			eachTask.TaskNode = task_node
			eachTask.TaskCommand = task_command
			this.Run(eachTask)
		} else {
			return
		}
	} else if btn_runAllTask != "" {
		if btn_taskDetail == "" {
			return
		} else {
			beego.Info(TaskJsonMap[btn_taskDetail])
			this.RunAllCmd(TaskJsonMap[btn_taskDetail])

		}
	}
	this.Redirect("/task", 302)
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

func (this *TaskController) Import() {
	filePath, err := this.FileUploadAndSave("importfiletasks", ".json")
	beego.Info(filePath)
	if err != nil {
		beego.Error(err)
	}
	TaskJsonMap, err = utils.TaskJsonRead(filePath)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(TaskJsonMap)
	// this.Data["taskData"] = TaskJsonMap
}

func (this *TaskController) Run(models.EachTask) {

}

func (this *TaskController) RunAllCmd(models.MainTasks) {

}
