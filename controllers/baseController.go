package controllers

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"path"
	"syscall"
	"time"

	"github.com/astaxie/beego"
)

//The Init Definition for main controller
type BaseController struct {
	beego.Controller
}

func (this *BaseController) FileUploadAndSave(inputname, fileformat string) (filePath string, err error) {
	File, FileReader, err := this.GetFile(inputname)
	if err != nil {
		beego.Error(err)
		return "", err
	}
	ext := path.Ext(FileReader.Filename)
	// check the fileformat
	var AllowExtMap map[string]bool = map[string]bool{
		fileformat: true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		beego.Error("we do not support to import this type file")
		return "", nil
	}
	// Create Directory
	uploadDir := "static/upload/" + time.Now().Format("2006-01-02")
	oldMask := syscall.Umask(0)
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		beego.Error(err)
		return "", err
	}
	syscall.Umask(oldMask)
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext
	filePath = uploadDir + "/" + fileName
	defer File.Close()
	err = this.SaveToFile(inputname, filePath)
	if err != nil {
		beego.Error(err)
		return "", err
	}
	return filePath, nil
}
