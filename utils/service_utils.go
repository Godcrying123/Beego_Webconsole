package utils

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
	"webconsole_sma/models"

	"github.com/astaxie/beego"
)

func init() {}

func Service_List() {
	// check all services in this machine

}

func Services_XMLGenerator(services map[string]models.Service) (message string, err error) {
	//err_chan := make(chan error)
	//output_chan := make(chan []byte, 3)
	//var wg sync.WaitGroup
	var builder strings.Builder
	var index int64
	var missrecord int
	builder.Write([]byte(xml.Header))
	for _, service_entity := range services {
		if service_entity.ServiceName == "" {
			missrecord++
			continue
		} else {
			service_entity.ID = index
			index++
			service_entity.LastStatusModifiedTime = time.Now()
			output, err := xml.MarshalIndent(service_entity, "", "\t")
			if err != nil {
				beego.Error(err)
				return "", err
			}
			//os.Stdout.Write(output)
			builder.Write(output)
			builder.Write([]byte("\n"))
		}

	}
	xmlOutPutDate := builder.String()
	filepath := "xmls/requirements_services.xml"
	ioutil.WriteFile(filepath, []byte(xmlOutPutDate), os.ModeAppend)
	if missrecord != 0 {
		return strconv.Itoa(missrecord) + " records cannot be exported, please check if there are an empty line in the list", nil
	}
	return "all services has been exported to XML successfully!", nil
}
