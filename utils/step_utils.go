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

func init() {

}

func StepAnalyzer(mainsteplist, stepname, stepsummary, stepcommand []string) map[string]models.MainSteps {
	var mainsteps models.MainSteps
	var substep models.SubSteps
	var substepslist = []models.SubSteps{}
	var mainstepmap = make(map[string]models.MainSteps)
	mainstepindex := 0
	listlength := len(mainsteplist)
	previousmainsteptitle := mainsteplist[0]
	for index := 0; index < listlength; index++ {
		mainsteps.StepTitle = mainsteplist[index]
		if mainsteplist[index] == previousmainsteptitle {
			substep = substepdefinition(stepname[index], stepsummary[index], stepcommand[index], index)
			substepslist = append(substepslist, substep)
		} else {
			mainsteps.StepTitle = previousmainsteptitle
			mainsteps.SubSteps = substepslist
			mainsteps.ID = mainstepindex
			mainstepindex++
			mainstepmap[previousmainsteptitle] = mainsteps
			substepslist = []models.SubSteps{}
			previousmainsteptitle = mainsteplist[index]
			substep = substepdefinition(stepname[index], stepsummary[index], stepcommand[index], index)
			substepslist = append(substepslist, substep)
		}
	}
	if mainsteps.StepTitle == mainsteplist[listlength-1] {
		mainsteps.SubSteps = substepslist
		mainsteps.ID = mainstepindex
		mainstepmap[previousmainsteptitle] = mainsteps
	} else {
		mainsteps.StepTitle = mainsteplist[listlength-1]
		mainsteps.SubSteps = substepslist
		mainsteps.ID = mainstepindex
		mainstepmap[previousmainsteptitle] = mainsteps
	}
	return mainstepmap
}

func substepdefinition(stepname, stepsummary, stepcommand string, index int) (substepreturn models.SubSteps) {
	substepreturn.StepID = index
	substepreturn.StepName = stepname
	substepreturn.StepSummary = stepsummary
	substepreturn.StepCommand = stepcommand
	return substepreturn
}

func StepJsonGenerator(mapstepsmap map[string]models.MainSteps) (message string, err error) {
	var builder strings.Builder
	for _, stepentity := range mapstepsmap {
		output, err := json.MarshalIndent(stepentity, "", "\t")
		if err != nil {
			beego.Error(err)
			return "", err
		}
		builder.Write(output)
		builder.WriteString(",\n")
	}
	err = WriteJson(builder, "json/requirements_steps.json")
	return "all steps have been exported to JSON successfully!", nil
}

func StepJsonRead(filePath string) (jsonStruct map[string]models.MainSteps, err error) {
	var byter bytes.Buffer
	jsonFile, err := ioutil.ReadFile(filePath)
	jsonStruct1 := make(map[string]models.MainSteps)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	byter.Write([]byte("["))
	byter.Write(jsonFile)
	byter.Write([]byte("]"))
	jsons, _ := simplejson.NewJson(byter.Bytes())
	// beego.Info(jsons)
	for _, jsonmap := range jsons.MustArray() {
		step := models.MainSteps{}
		err = mapstructure.WeakDecode(jsonmap.(map[string]interface{}), &step)
		if err != nil {
			beego.Error(err)
		}
		jsonStruct1[step.StepTitle] = step
	}
	beego.Info(jsonStruct1)
	return jsonStruct1, nil
}
