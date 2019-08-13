package controllers

import (
	"container/list"
	"os"
	"path"
	"webconsole_sma/models"

	"github.com/astaxie/beego"
)

var root_folder string

const fs_maxbufsize = 4096

//The Init Definition for main controller
type FileController struct {
	BaseController
}

func (this *FileController) Get() {
	this.TplName = "file.html"
	beego.Info(this.Ctx.Request.RequestURI)
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
	root_folder = path
	this.handleFile()
	beego.Info("All Files Infos have been listed successfully")
	return nil

}

func (this *FileController) handleFile() (err error) {
	urlstring := this.Ctx.Request.RequestURI
	filepath := path.Join((root_folder), urlstring[5:])
	beego.Info(filepath)
	err = this.serveFile(filepath)
	if err != nil {
		beego.Error(err)
	}
	return nil
}

func (this *FileController) serveFile(filepath string) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer file.Close()
	statinfo, err := file.Stat()
	if err != nil {
		beego.Error(err)
		return err
	}
	if statinfo.IsDir() {
		this.handleDirectory(file)
		return
	}
	if (statinfo.Mode() &^ 07777) == os.ModeSocket {
		beego.Info("403 Forbidden : you can't access this resource.")
		return
	}
	return nil
}

func copyToArray(src *list.List) []string {
	dst := make([]string, src.Len())

	i := 0
	for e := src.Front(); e != nil; e = e.Next() {
		dst[i] = e.Value.(string)
		i = i + 1
	}

	return dst
}

func (this *FileController) handleDirectory(file *os.File) {
	urlstring := this.Ctx.Request.RequestURI
	names, _ := file.Readdir(-1)
	for _, val := range names {
		if val.Name() == "index.html" {
			this.serveFile(path.Join(file.Name(), "index.html"))
			return
		}
	}

	// Otherwise, generate folder content.
	children_dir_tmp := list.New()
	children_files_tmp := list.New()

	for _, val := range names {
		if val.Name()[0] == '.' {
			continue
		}
		if val.IsDir() {
			children_dir_tmp.PushBack(val.Name())
		} else {
			children_files_tmp.PushBack(val.Name())
		}
	}

	// And transfer the content to the final array structure
	children_dir := copyToArray(children_dir_tmp)
	children_files := copyToArray(children_files_tmp)

	fileData := models.DirListing{
		Name:           urlstring[5:],
		Children_dir:   children_dir,
		Children_files: children_files,
	}

	this.Data["fileData"] = fileData

	beego.Info(fileData)
}
