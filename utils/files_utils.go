package utils

import (
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/astaxie/beego"
)

func init() {}

func WriteJson(builder strings.Builder, filepath string) error {
	jsonOutPutDate := strings.TrimRight(builder.String(), ",\n")
	//os.Stdout.Write([]byte(jsonOutPutDate))
	oldMask := syscall.Umask(0)
	err := ioutil.WriteFile(filepath, []byte(jsonOutPutDate), os.ModeAppend)
	if err != nil {
		beego.Error(err)
	}
	syscall.Umask(oldMask)
	return err
}
