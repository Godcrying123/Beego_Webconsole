package controllers

import (
	"os"
	"strings"
	"webconsole_sma/models"
	"webconsole_sma/utils"

	"golang.org/x/crypto/ssh"

	"github.com/astaxie/beego"
	"github.com/pkg/sftp"
)

var (
	sftpPath  string
	sshConn   *ssh.Client
	sftpConn  *sftp.Client
	BaseUrl   string
	navurltmp *string
)

type STFPController struct {
	beego.Controller
}

func (this *STFPController) Get() {
	this.TplName = "file.html"
	this.Data["stepsData"] = StepJsonStruct
	this.Data["services"] = JsonStruct
	urlstring = this.Ctx.Request.RequestURI
	navurl = strings.Split(urlstring, "/file/")
	sftpPath = "/" + navurl[1]
	BaseUrl = navurl[0] + "/"
	HostName = this.Ctx.Input.Param(":hostname")
	sshHost := SSHHosts[HostName]
	filename := this.Input().Get("editfile")
	dlfilename := this.Input().Get("dl")
	newSSHRemoteAddr := sshHost.HostIp + ":" + sshHost.SSHPort
	if sshConn == nil || sshConn.Conn.RemoteAddr().String() != newSSHRemoteAddr {
		var err error
		sshConn, err = utils.NewSshClient(sshHost)
		if err != nil {
			beego.Error(err)
		}
		sftpConn, err = sftp.NewClient(sshConn)
		if err != nil {
			beego.Error(err)
		}
		if filename != "" || dlfilename != "" {
			err = this.editAndDownFiles(filename, dlfilename, sftpConn)
			if err != nil {
				beego.Error(err)
			}
		}
		err = this.FileList(sftpConn)
		if err != nil {
			beego.Error(err)
		}
	} else {
		if filename != "" || dlfilename != "" {
			err := this.editAndDownFiles(filename, dlfilename, sftpConn)
			if err != nil {
				beego.Error(err)
			}
		}
		err := this.FileList(sftpConn)
		if err != nil {
			beego.Error(err)
		}
	}
	this.Data["navUrl"] = navurlsmap
	// defer sshConn.Close()
	// defer sftpConn.Close()
}

func (this *STFPController) editAndDownFiles(fileName, dlFileName string, sftpConn *sftp.Client) (err error) {
	navurltmp := editAndSFTPURLParse()
	sftpPath = navurltmp[0]
	if fileName != "" {
		err = this.editAndReadFiles(fileName, sftpConn)
		if err != nil {
			beego.Error(err)
			return err
		}
	} else if dlFileName != "" {
		err = this.downloadFiles(dlFileName, sftpConn)
		if err != nil {
			beego.Error(err)
			return err
		}
	}
	return nil
}

func (this *STFPController) editAndReadFiles(fileName string, sftpConn *sftp.Client) (err error) {

	return nil
}

func (this *STFPController) downloadFiles(dlFileName string, sftpConn *sftp.Client) (err error) {
	downloadFile, err := sftpConn.OpenFile(sftpPath+dlFileName, os.O_WRONLY)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer downloadFile.Close()
	statinfo, err := downloadFile.Stat()
	if err != nil {
		return err
	}
	if statinfo.IsDir() {

	} else {
		dstFile, _ := os.Create("/tmp/test.txt")
		defer dstFile.Close()
		srcFile, _ := sftpConn.Open("./file.txt")
		defer srcFile.Close()
	}
	return nil
}

func editAndSFTPURLParse() []string {
	navurltmp := strings.Split(sftpPath, "?")
	urlstring = navurltmp[0]
	navurl = strings.Split(navurltmp[0][5:], "/")
	return navurltmp
}

func (this *STFPController) FileList(sftpConn *sftp.Client) (err error) {
	_, err = sftpConn.Stat(sftpPath)
	if os.IsNotExist(err) {
		beego.Error(err)
		return err
	}
	err = this.handleFile(sftpConn)
	if err != nil {
		beego.Error(err)
		return err
	}
	beego.Info("All Files Infos have been listed successfully")
	return nil
}

func (this *STFPController) handleFile(sftpConn *sftp.Client) (err error) {
	_, err = sftpConn.Stat(sftpPath)
	if err != nil {
		beego.Error(err)
		return err
	}
	walkFiles, err := sftpConn.ReadDir(sftpPath)
	if err != nil {
		beego.Error(err)
		return err
	}
	var childrenDirTmp models.Directory
	var childrenDirs []models.Directory
	var childrenFilesTmp models.File
	var childrenFiles []models.File
	for _, listFile := range walkFiles {
		// if listFile.Name()[0] == "."
		if listFile.IsDir() {
			childrenDirTmp.DirName = listFile.Name()
			childrenDirTmp.DirSize = listFile.Size()
			childrenDirTmp.DirLastModified = listFile.ModTime()
			childrenDirTmp.DirAccess = listFile.Mode()
			childrenDirTmp.DirPath = BaseUrl + "file" + sftpPath
			childrenDirs = append(childrenDirs, childrenDirTmp)
		} else {
			childrenFilesTmp.FileName = listFile.Name()
			childrenFilesTmp.FileSize = listFile.Size()
			childrenFilesTmp.FileLastModified = listFile.ModTime()
			childrenFilesTmp.FileAccess = listFile.Mode()
			childrenFilesTmp.FilePath = BaseUrl + "file" + sftpPath
			childrenFiles = append(childrenFiles, childrenFilesTmp)
		}

		sftSFTPData := models.DirListing1{
			Name:          sftpPath,
			ChildrenDirs:  childrenDirs,
			ChildrenFiles: childrenFiles,
		}
		this.Data["fileData"] = sftSFTPData
	}
	return nil
}
