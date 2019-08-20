package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"
	"webconsole_sma/models"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"github.com/mitchellh/mapstructure"
)

type ServiceSlice struct {
	Services []models.Service `json:"service"`
}

func init() {}

func ServicesJsonGenerator(services map[string]models.Service) (message string, err error) {
	//err_chan := make(chan error)
	//output_chan := make(chan []byte, 3)
	//var wg sync.WaitGroup
	var builder strings.Builder
	var index int64
	var missrecord int
	for _, service_entity := range services {
		if service_entity.ServiceName == "" {
			missrecord++
			continue
		} else {
			service_entity.ID = index
			index++
			output, err := json.MarshalIndent(service_entity, "", "\t")
			if err != nil {
				beego.Error(err)
				return "", err
			}
			builder.Write(output)
			builder.WriteString(",\n")
		}
	}
	err = WriteJson(builder, "json/requirements_services.json")
	if missrecord != 0 {
		return strconv.Itoa(missrecord) + " records cannot be exported, please check if there are an empty line in the list", nil
	}
	return "all services has been exported to JSON successfully!", nil
}

func ServicesJsonRead(filePath string) (jsonStruct map[string]models.Service, err error) {
	var byter bytes.Buffer
	jsonFile, err := ioutil.ReadFile(filePath)
	jsonStruct1 := make(map[string]models.Service)
	if err != nil {
		beego.Error(err)
		return nil, err
	}
	byter.Write([]byte("["))
	byter.Write(jsonFile)
	byter.Write([]byte("]"))
	jsons, _ := simplejson.NewJson(byter.Bytes())
	for _, jsonmap := range jsons.MustArray() {
		service := models.Service{}
		err = mapstructure.WeakDecode(jsonmap.(map[string]interface{}), &service)
		if err != nil {
			beego.Error(err)
		}
		jsonStruct1[service.ServiceName] = service
	}
	return jsonStruct1, nil
}

func ServiceInfo(Service models.Service) (ServiceReturn models.Service, err error) {
	ServiceReturn = Service
	activeAndrun, err := ServiceStatus(ServiceReturn.ServiceName)
	if err != nil {
		beego.Error(err)
	}
	statusList := strings.Split(activeAndrun, " ")
	if len(statusList) >= 2 {
		ServiceReturn.ActiveStatus = strings.Trim(statusList[0], " ")
		ServiceReturn.RunningStatus = strings.Trim(statusList[1], "\n")
	}
	ServiceReturn.ServiceStatus, err = ServiceDetail(ServiceReturn.ServiceName)
	if err != nil {
		beego.Error(err)
		return ServiceReturn, err
	}
	return ServiceReturn, nil
}

func ServiceStatus(servicename string) (string, error) {
	var builder strings.Builder
	builder.WriteString("systemctl -a |grep ")
	builder.WriteString("'")
	builder.WriteString(servicename)
	builder.WriteString(".service'|awk '{print $3 FS $4}'")
	servicestatus, err := CommandExecReturnString(builder.String())
	if err != nil {
		beego.Error(err)
		return "", err
	}
	return servicestatus, nil
}

func ServiceDetail(servicename string) (string, error) {
	var builder strings.Builder
	builder.Write([]byte("systemctl status -l "))
	builder.Write([]byte(servicename))
	builder.Write([]byte(".service"))
	servicedetail, err := CommandExecReturnString(builder.String())
	if err != nil {
		beego.Error(err)
		return "", err
	}
	return servicedetail, nil
}
