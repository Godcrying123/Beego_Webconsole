package controllers

import (
	"os"

	"github.com/astaxie/beego"
)

var root_folder *string

const fs_maxbufsize = 4096

//The Init Definition for main controller
type FileController struct {
	BaseController
}

func (this *FileController) Get() {
	this.TplName = "file.html"
	this.FileList("/")
}

func (this *FileController) FileList(path string) error {
	beego.Info("Starting to share the File List")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		beego.Error(err)
		return err
	}
	// root_folder = flag.String("root", path, "Root folder")
	// flag.Parse()
	// handleFile()
	beego.Info("All Files Infos have been listed successfully")
	return nil

}

func handleFile() {

}
