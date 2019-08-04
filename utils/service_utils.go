package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"webconsole_sma/models"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

type ServiceSlice struct {
	Services []models.Service
}

func init() {}

func Services_JsonGenerator(services map[string]models.Service) (message string, err error) {
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
			service_entity.LastStatusModifiedTime = time.Now()
			output, err := json.MarshalIndent(service_entity, "", "\t")
			if err != nil {
				beego.Error(err)
				return "", err
			}
			builder.Write(output)
			builder.Write([]byte("\n"))
		}

	}
	jsonOutPutDate := builder.String()
	os.Stdout.Write([]byte(jsonOutPutDate))
	oldMask := syscall.Umask(0)
	filepath := "json/requirements_services.json"
	err = ioutil.WriteFile(filepath, []byte(jsonOutPutDate), os.ModeAppend)
	if err != nil {
		beego.Error(err)
	}
	syscall.Umask(oldMask)
	if missrecord != 0 {
		return strconv.Itoa(missrecord) + " records cannot be exported, please check if there are an empty line in the list", nil
	}
	return "all services has been exported to JSON successfully!", nil
}

func Services_JsonRead(filePath string) (jsonStruct map[string]models.Service, err error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		beego.Error(err)
	}
	//var services_list ServiceSlice
	jsonString, err := Json_to_byte(jsonFile)
	defer jsonFile.Close()
	if err != nil {
		beego.Error(err)
	}
	res, err := simplejson.NewJson([]byte(jsonString))
	if err != nil {
		beego.Error(err)
	} else {
		servicename_list, _ := res.Get("ServiceName").StringArray()
		beego.Info(servicename_list)
		//serviceversion_list, err := res.Get("ServiceVersion").StringArray()
		for name := range servicename_list {
			beego.Info(name)
		}
	}
	return
}

func Service_List() {
	// check all services in this machine

}

func Json_to_byte(file *os.File) (string, error) {
	var builder strings.Builder
Loop:
	for {
		buf := make([]byte, 1024)
		switch nr, err := file.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "Json_to_byte: error reading %s\n", err.Error())
			return "", err
		case nr > 0:
			builder.Write(buf)
		case nr == 0:
			break Loop
		}
	}
	return builder.String(), nil
}
