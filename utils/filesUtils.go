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

func FileWrite(File models.File) (err error) {
	// oldMask := syscall.Umask(0)
	file, err := os.OpenFile(File.FilePath+File.FileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 766)
	// syscall.Umask(oldMask)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(File.FileContent); err != nil {
		beego.Error(err)
		return err
	}
	return nil
}

func FileExistCheck(fileNameAndPath string) (Exist bool, err error) {
	_, err = os.Stat(fileNameAndPath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
