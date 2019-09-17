package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strings"
	"webconsole_sma/models"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"github.com/mitchellh/mapstructure"
)

func TaskAnalyzer(taskname, tasksummary, tasknode, taskcommand []string) (maintaskslices []models.MainTasks) {
	var tasksteps models.MainTasks
	var eachtask models.EachTask
	var tasklist = []models.EachTask{}
	var tasklistslice = []models.MainTasks{}
	taskindex := 0
	listlength := len(taskname)
	previoustasktitle := taskname[0]
	for index := 0; index < listlength; index++ {
		tasksteps.TaskTitle = taskname[index]
		if taskname[index] == previoustasktitle {
			eachtask = taskdefinition(tasksummary[index], tasknode[index], taskcommand[index], index)
			tasklist = append(tasklist, eachtask)
		} else {
			tasksteps.TaskTitle = previoustasktitle
			tasksteps.SubTasks = tasklist
			tasksteps.ID = taskindex
			taskindex++
			tasklistslice = append(tasklistslice, tasksteps)
			tasklist = []models.EachTask{}
			previoustasktitle = taskname[index]
			eachtask = taskdefinition(tasksummary[index], tasknode[index], taskcommand[index], index)
			tasklist = append(tasklist, eachtask)
		}
	}
	if tasksteps.TaskTitle == taskname[listlength-1] {
		tasksteps.SubTasks = tasklist
		tasksteps.ID = taskindex
		tasklistslice = append(tasklistslice, tasksteps)
	} else {
		tasksteps.TaskTitle = taskname[listlength-1]
		tasksteps.SubTasks = tasklist
		tasksteps.ID = taskindex
		tasklistslice = append(tasklistslice, tasksteps)
	}
	return tasklistslice
}

func taskdefinition(tasksummary, tasknode, taskcommand string, index int) (taskreturn models.EachTask) {
	taskreturn = models.EachTask{
		TaskID:      index,
		TaskSummary: tasksummary,
		TaskNode:    tasknode,
		TaskCommand: taskcommand,
	}
	return taskreturn
}

func TaskJsonGenerator(tasklistslice []models.MainTasks) (message string, err error) {
	var builder strings.Builder
	for _, taskentity := range tasklistslice {
		output, err := json.MarshalIndent(taskentity, "", "\t")
		if err != nil {
			beego.Error(err)
			return "", err
		}
		builder.Write(output)
		builder.WriteString(",\n")
	}
	err = WriteJson(builder, "json/automation_task.json")
	return "all tasks have been exported to JSON successfully!", nil
}

func TaskJsonRead(filePath string) (map[string]models.MainTasks, error) {
	var byter bytes.Buffer
	jsonFile, err := ioutil.ReadFile(filePath)
	var taskjsonMap = make(map[string]models.MainTasks)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	byter.Write([]byte("["))
	byter.Write(jsonFile)
	byter.Write([]byte("]"))
	jsons, _ := simplejson.NewJson(byter.Bytes())
	for _, jsonmap := range jsons.MustArray() {
		task := models.MainTasks{}
		err = mapstructure.WeakDecode(jsonmap.(map[string]interface{}), &task)
		if err != nil {
			beego.Error(err)
		}
		taskjsonMap[task.TaskTitle] = task
	}
	return taskjsonMap, nil
}

func SSHConnTaskRun(node models.MachineSSH, cmd string) (err error) {
	beego.Info(cmd)
	sshConn, err := NewSshClient(node)
	if err != nil {
		return err
	}
	defer sshConn.Close()
	sshSess, err := sshConn.NewSession()
	if err != nil {
		return err
	}
	defer sshSess.Close()
	buf, err := sshSess.CombinedOutput(cmd)
	if err != nil {
		return err
	}
	beego.Info(string(buf))
	return nil
	// sshSess.Stdout = os.Stdout
	// sshSess.Stderr = os.Stderr
	// err = sshSess.Shell()
	// if err != nil {
	// 	return err
	// }
	// for _, cmd := range cmd {

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// _, err = fmt.Fprintf(stdin, "%s\n", cmd)
	// b, err := sshSess.Output()
	// out := string(b)
	// err = sshSess.Wait()
}
