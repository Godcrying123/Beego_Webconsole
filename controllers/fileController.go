package controllers

import (
	"os"
	"path"
	"strings"
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"github.com/astaxie/beego"
)

var (
	root_folder  string
	navurlstring string
	urlstring    string
	navurl       []string
	navurlsmap   map[string]string
	File         models.File
	uses_gzip    *bool
)

const fs_maxbufsize = 4096

//The Init Definition for main controller
type FileController struct {
	beego.Controller
}

func (this *FileController) Get() {
	this.TplName = "file.html"
	this.Data["stepsData"] = StepJsonStruct
	this.Data["services"] = JsonStruct
	filename := this.Input().Get("editfile")
	dlfilename := this.Input().Get("dl")
	navurlsmap = make(map[string]string)
	navurlsmap["/"] = "/"
	urltmp := ""
	urlstring = this.Ctx.Request.RequestURI
	if filename != "" {
		navurltmp := editAndDLUrlParse()
		File, err := utils.FileRead(filename, navurltmp[0][5:])
		if err != nil {
			beego.Error(err)
		}
		this.Data["File"] = File
	} else if dlfilename != "" {
		navurltmp := editAndDLUrlParse()
		beego.Info(navurltmp[0][5:])
		beego.Info(dlfilename)
		err := utils.DirAndFileDownLoad(this.Ctx.ResponseWriter, this.Ctx.Request, navurltmp[0][5:], dlfilename)
		if err != nil {
			beego.Error(err)
		}

	} else {
		this.FileList("/")
		navurl = strings.Split(urlstring[5:], "/")
	}
	this.FileList("/")
	this.Data["baseUrl"] = urlstring[5:]
	func() {
		for i := 1; i < len(navurl)-1; i++ {
			urltmp = urltmp + navurl[i] + "/"
			navurlsmap[navurl[i]] = urltmp
		}
	}()
	this.Data["navUrl"] = navurlsmap
}

func editAndDLUrlParse() []string {
	navurltmp := strings.Split(urlstring, "?")
	urlstring = navurltmp[0]
	navurl = strings.Split(navurltmp[0][5:], "/")
	return navurltmp
}

func (this *FileController) Post() {
	this.TplName = "file.html"
	btn_save := this.Input().Get("savefile")
	btn_find := this.Input().Get("findfile")
	filePath := this.Input().Get("savefilepath")
	fileName := this.Input().Get("savefilename")
	if btn_save != "" {
		fileContent := this.Input().Get("filecontent")
		// beego.Info(fileContent)
		// beego.Info(filePath)
		// beego.Info(fileName)
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
			beego.Info(urlstring)
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
	} else if btn_find != "" {
		if fileName == "" {
			urlstring = "/file" + filePath
		} else {
			_, err := utils.FileRead(fileName, filePath)
			if err != nil {
				urlstring = "/file" + filePath
				beego.Error(err)
			} else {
				urlstring = "/file" + filePath + "?editfile=" + fileName
			}
		}
		beego.Info(urlstring)
	}
	this.Redirect(urlstring, 302)
}

func (this *FileController) FileList(path string) error {
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
	func() {
		for _, val := range names {
			if val.Name()[0] == '.' {
				continue
			}
			if val.IsDir() {
				func() {
					childrenDirTmp.DirName = val.Name()
					childrenDirTmp.DirSize = val.Size()
					childrenDirTmp.DirLastModified = val.ModTime()
					childrenDirTmp.DirAccess = val.Mode()
					childrenDirTmp.DirPath = urlstring
					childrenDirs = append(childrenDirs, childrenDirTmp)
				}()
			} else {
				func() {
					childrenFilesTmp.FileName = val.Name()
					childrenFilesTmp.FileSize = val.Size()
					childrenFilesTmp.FileLastModified = val.ModTime()
					childrenFilesTmp.FileAccess = val.Mode()
					childrenFilesTmp.FilePath = urlstring
					childrenFiles = append(childrenFiles, childrenFilesTmp)
				}()
			}
		}
	}()
	// beego.Info(childrenDirs)
	// beego.Info(childrenFiles)
	fileData := models.DirListing{
		Name:          urlstring[5:],
		ChildrenDirs:  childrenDirs,
		ChildrenFiles: childrenFiles,
	}

	this.Data["fileData"] = fileData
}

func (this *FileController) GetSFTP() {
	this.TplName = "file.html"

}
