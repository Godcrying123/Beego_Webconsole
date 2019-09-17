package controllers

import (
	"strings"
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var (
	StepJsonStruct []models.MainSteps
	JsonStruct     map[string]models.Service
	TaskJsonMap    map[string]models.MainTasks
	ServiceClients = make(map[*websocket.Conn]bool)
	Servicechan    = make(chan models.Service)
	HostName       string
	SSHUrl         string
)

func init() {
	var err error
	JsonStruct, err = utils.ServicesJsonRead("json/requirements_services.json")
	if err != nil {
		beego.Error(err)
	}
	for servicekey, service := range JsonStruct {
		serviceReturn, err := utils.ServiceInfo(service)
		if err != nil {
			beego.Error(err)
		}
		JsonStruct[servicekey] = serviceReturn

	}
	StepJsonStruct, err = utils.StepJsonRead("json/requirements_steps.json")
	if err != nil {
		beego.Error(err)
	}
	TaskJsonMap, err = utils.TaskJsonRead("json/automation_task.json")
	if err != nil {
		beego.Error(err)
	}
}

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.TplName = "index.html"
	this.Data["stepsData"] = StepJsonStruct
	this.Data["services"] = JsonStruct
	this.Data["taskData"] = TaskJsonMap
	HostName = this.Ctx.Input.Param(":hostname")
	if HostName != "" {
		SSHUrl = "/node/" + HostName + "/"
	} else {
		SSHUrl = "/"
	}
	if HostName == "" {
		HostName = "localhost"
	}
	this.Data["sshUrl"] = SSHUrl
	this.Data["machine"] = SSHHosts
	// beego.Info(HostName)
}

func (this *IndexController) Post() {
	this.TplName = "index.html"
	btn_stepimport := this.Input().Get("importallsteps")
	btn_serviceimport := this.Input().Get("importall")
	btn_hostonoff := this.Input().Get("syncoff")
	btn_importTask := this.Input().Get("importalltasks")
	btn_runTask := this.Input().Get("runTask")
	btn_runAllTask := this.Input().Get("runAllTask")
	btn_taskDetail := this.Input().Get("AllTaskDetails")
	task_node := this.Input().Get("TaskNode")
	task_command := this.Input().Get("TaskCommand")
	if btn_stepimport != "" {
		this.StepImport()
	} else if btn_serviceimport != "" {
		this.ServiceImport()
	} else if btn_hostonoff != "" {
		hostOnAndOff = !hostOnAndOff
	} else if btn_importTask != "" {
		this.TaskImport()
	} else if btn_runTask != "" {
		eachTask := models.EachTask{}
		if task_node != "" && task_command != "" {
			eachTask.TaskNode = task_node
			eachTask.TaskCommand = task_command
			err := this.Run(eachTask)
			if err != nil {
				beego.Error(err)
			}
		} else {
			return
		}
	} else if btn_runAllTask != "" {
		if btn_taskDetail != "" {
			err := this.RunAllCmd(TaskJsonMap[btn_taskDetail])
			if err != nil {
				beego.Error(err)
			}
		} else {
			return
		}
	}
	redirectURL := strings.TrimRight(SSHUrl, "/")
	// beego.Info(redirectURL)
	this.Redirect(redirectURL, 301)
}

func (this *IndexController) TaskImport() {
	filePath, err := this.FileUploadAndSave("importfiletasks", ".json")
	if err != nil {
		beego.Error(err)
	}
	TaskJsonMap, err = utils.TaskJsonRead(filePath)
	if err != nil {
		beego.Error(err)
	}
	// this.Data["taskData"] = TaskJsonMap
}

func (this *IndexController) StepImport() {
	filePath, err := this.FileUploadAndSave("importfilestep", ".json")
	if err != nil {
		beego.Error(err)
		return
	}
	StepJsonStruct, err = utils.StepJsonRead(filePath)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 302)
}

func (this *IndexController) ServiceImport() {
	filePath, err := this.FileUploadAndSave("importfile", ".json")
	if err != nil {
		beego.Error(err)
	}
	JsonStruct, err = utils.ServicesJsonRead(filePath)

	if err != nil {
		beego.Error(err)
	}
	for servicekey, service := range JsonStruct {
		serviceReturn, err := utils.ServiceInfo(service)
		if err != nil {
			beego.Error(err)
		}
		JsonStruct[servicekey] = serviceReturn

	}
	this.Redirect("/", 302)
}

func (this *IndexController) Run(eachTask models.EachTask) (err error) {
	beego.Info("Starting to run command")
	err = utils.SSHConnTaskRun(SSHHosts[eachTask.TaskNode], eachTask.TaskCommand)
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}

func (this *IndexController) RunAllCmd(mainTask models.MainTasks) (err error) {
	for _, taskentity := range mainTask.SubTasks {
		beego.Info(taskentity.TaskSummary)
		err := utils.SSHConnTaskRun(SSHHosts[taskentity.TaskNode], taskentity.TaskCommand)
		if err != nil {
			beego.Error(err)
			return err
		}
	}
	return nil
}
