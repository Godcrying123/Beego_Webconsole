package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"webconsole_sma/models"

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

func FileRead(filename, filepath string) (File models.File, err error) {
	var builder strings.Builder
	file, err := os.Open(filepath + filename)
	File.FileName = filename
	File.FilePath = filepath
	if err != nil {
		return File, err
	}
	defer file.Close()
	bufReader := bufio.NewReader(file)
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		} else {
			builder.Write(line)
			builder.WriteString("\n")
		}
	}
	fileString := builder.String()
	File.FileContent = fileString
	return File, err
}
