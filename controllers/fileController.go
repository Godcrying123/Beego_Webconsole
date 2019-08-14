package controllers

import (
	"os"
	"path"
	"strings"
	"webconsole_sma/models"

	"github.com/astaxie/beego"
)

var root_folder string
var navurlstring string
var urlstring string
var navurl []string
var navurls []string

const fs_maxbufsize = 4096

//The Init Definition for main controller
type FileController struct {
	beego.Controller
}

func (this *FileController) Get() {
	this.TplName = "fileTable.html"
	urlstring = this.Ctx.Request.RequestURI
	this.FileList("/")
	navurl = strings.Split(urlstring, "/")
	navurlnospace := []string{}
	navurls = make([]string, len(navurl)+1)
	navurls[0] = "/"
	for urlindex := 0; urlindex < len(navurl); urlindex++ {
		if strings.TrimSpace(navurl[urlindex]) != "" {
			navurlnospace = append(navurlnospace, navurl[urlindex])
		} else {
			continue
		}
	}
	for urlindex := 0; urlindex < len(navurlnospace); urlindex++ {
		urltmp := navurls[urlindex] + navurlnospace[urlindex] + "/"
		navurls[urlindex+1] = urltmp
	}

	beego.Info(navurls)
	this.Data["navUrl"] = navurls
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

func (this *FileController) handleDirectory(file *os.File) {
	names, _ := file.Readdir(-1)
	for _, val := range names {
		if val.Name() == "index.html" {
			this.serveFile(path.Join(file.Name(), "index.html"))
			return
		}
	}

	// Otherwise, generate folder content.
	var childrenDirTmp models.Directory
	var childrenDirs []models.Directory
	var childrenFilesTmp models.File
	var childrenFiles []models.File

	for _, val := range names {
		if val.Name()[0] == '.' {
			continue
		}
		if val.IsDir() {
			childrenDirTmp.DirName = val.Name()
			childrenDirTmp.DirSize = val.Size()
			childrenDirTmp.DirLastModified = val.ModTime()
			childrenDirTmp.DirAccess = val.Mode()
			childrenDirTmp.DirPath = urlstring[5:]
			childrenDirs = append(childrenDirs, childrenDirTmp)
		} else {
			childrenFilesTmp.FileName = val.Name()
			childrenFilesTmp.FileSize = val.Size()
			childrenFilesTmp.FileLastModified = val.ModTime()
			childrenFilesTmp.FileAccess = val.Mode()
			childrenFilesTmp.FilePath = urlstring[5:]
			childrenFiles = append(childrenFiles, childrenFilesTmp)
		}
	}

	// beego.Info(childrenDirs)
	// beego.Info(childrenFiles)
	fileData := models.DirListing1{
		Name:          urlstring[5:],
		ChildrenDirs:  childrenDirs,
		ChildrenFiles: childrenFiles,
	}

	this.Data["fileData"] = fileData
}
