package controllers

import (
	"os"
	"path"
	"strings"
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

var root_folder string
var navurlstring string
var urlstring string
var navurl []string
var navurlsmap map[string]string
var File models.File

const fs_maxbufsize = 4096

//The Init Definition for main controller
type FileController struct {
	beego.Controller
}

func (this *FileController) Get() {
	this.TplName = "file.html"
	filename := this.Input().Get("editfile")
	navurlsmap = make(map[string]string)
	navurlsmap["/"] = "/"
	urltmp := ""
	urlstring = this.Ctx.Request.RequestURI
	if filename != "" {
		navurltmp := strings.Split(urlstring, "?")
		urlstring = navurltmp[0]
		navurl = strings.Split(navurltmp[0][5:], "/")
		File, err := utils.FileRead(filename, navurltmp[0][5:])
		if err != nil {
			beego.Error(err)
		}
		this.Data["File"] = File
		this.Data["baseUrl"] = urlstring[5:]
		this.FileList("/")
	} else {
		this.FileList("/")
		navurl = strings.Split(urlstring[5:], "/")
		this.Data["baseUrl"] = urlstring[5:]
	}
	func() {
		for i := 1; i < len(navurl)-1; i++ {
			urltmp = urltmp + navurl[i] + "/"
			navurlsmap[navurl[i]] = urltmp
		}
	}()
	this.Data["navUrl"] = navurlsmap
}

func (this *FileController) Post() {
	this.TplName = "file.html"
	fileContent := this.Input().Get("filecontent")
	filePath := this.Input().Get("savefilepath")
	fileName := this.Input().Get("savefilename")
	beego.Info(fileContent)
	beego.Info(filePath)
	beego.Info(fileName)
	ok, _ := utils.FileExistCheck(filePath + fileName)
	if ok == false {
		beego.Info("I am here")
		beego.Info(ok)
		File.FileName = fileName
		File.FilePath = filePath
		File.FileContent = fileContent
		err := utils.FileWrite(File)
		if err != nil {
			beego.Error(err)
		}
	} else {
		beego.Info("I am here")
		File, err := utils.FileRead(fileName, filePath)
		if err != nil {
			beego.Error(err)
		}
		if File.FileContent != fileContent {
			File.FileContent = fileContent
		}
		err = utils.FileWrite(File)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Redirect(urlstring, 302)
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
