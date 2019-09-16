package controllers

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io"
	"net/http"
	"os"
	"strconv"
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
	this.Data["machine"] = SSHHosts
	this.Data["stepsData"] = StepJsonStruct
	this.Data["services"] = JsonStruct
	this.Data["sshUrl"] = SSHUrl
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
	this.Data["baseUrl"] = sftpPath
}

func (this *STFPController) Post() {
	this.TplName = "file.html"
	btn_save := this.Input().Get("savefile")
	btn_find := this.Input().Get("findfile")
	filePath := this.Input().Get("savefilepath")
	fileName := this.Input().Get("savefilename")
	// beego.Info(sftpPath)
	if btn_save != "" {
		fileContent := this.Input().Get("filecontent")
		sftpWriteFile, err := sftpConn.OpenFile(filePath+fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
		defer sftpWriteFile.Close()
		File.FileName = fileName
		File.FilePath = filePath
		File.FileContent = fileContent
		if err != nil {
			beego.Error(err)
		} else if os.IsNotExist(err) {
			beego.Info("This File Not Existed!")
			if _, err := sftpWriteFile.Write([]byte(File.FileContent)); err != nil {
				beego.Error(err)
			}
		} else {
			if _, err := sftpWriteFile.Write([]byte(File.FileContent)); err != nil {
				beego.Error(err)
			}
		}
		this.Redirect(BaseUrl+"file"+urlstring, 302)
	} else if btn_find != "" {
		if fileName == "" {
			urlstring = "/file" + filePath
		} else {
			_, err := sftpConn.Open(filePath + fileName)
			if err != nil {
				urlstring = "/file" + filePath
				beego.Error(err)
			} else {
				urlstring = "/file" + filePath + "?editfile=" + fileName
			}
		}
		this.Redirect(BaseUrl+urlstring, 302)
	}
}

func (this *STFPController) editAndDownFiles(fileName, dlFileName string, sftpConn *sftp.Client) (err error) {
	navurltmp := editAndSFTPURLParse()
	sftpPath = navurltmp[0]
	if fileName != "" {
		err = this.readFiles(fileName, sftpConn)
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

func (this *STFPController) readFiles(fileName string, sftpConn *sftp.Client) (err error) {
	var builder strings.Builder
	var showFile models.File
	showFile.FileName = fileName
	showFile.FilePath = sftpPath
	readFile, err := sftpConn.Open(sftpPath + fileName)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer readFile.Close()
	bufReader := bufio.NewReader(readFile)
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
	showFile.FileContent = fileString
	this.Data["File"] = showFile
	return nil
}

func (this *STFPController) downloadFiles(dlFileName string, sftpConn *sftp.Client) (err error) {
	downloadFile, err := sftpConn.Open(sftpPath + dlFileName)
	writer := this.Ctx.ResponseWriter
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
		zipFileName := strings.Trim(dlFileName, "/") + ".zip"
		err := compressFiles(sftpPath+dlFileName, sftpPath+zipFileName, sftpConn)
		if err != nil {
			beego.Error(err)
			return err
		}
		zipFile, err := sftpConn.Open(sftpPath + zipFileName)
		if err != nil {
			beego.Error(err)
			return err
		}
		defer zipFile.Close()
		err = dirAndFileDownloadToClient(writer, this.Ctx.Request, zipFile, zipFileName)
		if err != nil {
			beego.Error(err)
			return err
		}
		err = sftpConn.Remove(sftpPath + zipFileName)
		if err != nil {
			beego.Error(err)
			return err
		}
	} else {
		err = dirAndFileDownloadToClient(writer, this.Ctx.Request, downloadFile, dlFileName)
		if err != nil {
			beego.Error(err)
			return err
		}
	}
	defer writer.Flush()
	return nil
}

func dirAndFileDownloadToClient(writer http.ResponseWriter, request *http.Request, downloadFile *sftp.File, dlFileName string) (err error) {
	FileHeader := make([]byte, 512)
	downloadFile.Read(FileHeader)
	FileContentType := http.DetectContentType(FileHeader)
	//Get the file size
	statinfo, err := downloadFile.Stat()
	if err != nil {
		beego.Error(err)
		return err
	}
	FileSize := strconv.FormatInt(statinfo.Size(), 10) //Get file size as a string
	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+dlFileName)
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)
	downloadFile.Seek(0, 0)
	_, err = io.Copy(writer, downloadFile)
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}

func compressFiles(baseFolder, zipFilePath string, sftpConn *sftp.Client) (err error) {
	outFile, err := sftpConn.Create(zipFilePath)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)

	err = addFiles(zipWriter, baseFolder, "", sftpConn)
	if err != nil {
		beego.Error(err)
		return err
	}

	err = zipWriter.Close()
	if err != nil {
		beego.Error(err)
		return err
	}

	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string, sftpConn *sftp.Client) (err error) {
	files, err := sftpConn.ReadDir(basePath)
	if err != nil {
		beego.Error(err)
		return err
	}

	for _, file := range files {
		// beego.Info(basePath + file.Name())
		if !file.IsDir() {
			subFile, err := sftpConn.Open(basePath + file.Name())
			if err != nil {
				beego.Error(err)
				return err
			}
			defer subFile.Close()
			var n int64 = bytes.MinRead
			if fi, err := subFile.Stat(); err == nil {
				if size := fi.Size() + bytes.MinRead; size > n {
					n = size
				}
			}
			data, err := readAll(subFile, n)
			if err != nil {
				beego.Error(err)
				return err
			}
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				beego.Error(err)
				return err
			}
			_, err = f.Write(data)
			if err != nil {
				beego.Error(err)
				return err
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			// fmt.Println("Recursing and Adding SubDir: " + file.Name())
			// fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/", sftpConn)
		}
	}
	return nil
}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	var buf bytes.Buffer
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	if int64(int(capacity)) == capacity {
		buf.Grow(int(capacity))
	}
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

func editAndSFTPURLParse() []string {
	navurltmp := strings.Split(sftpPath, "?")
	// beego.Info(len(navurltmp))
	urlstring = navurltmp[0]
	// beego.Info(navurltmp)
	// navurl = strings.Split(navurltmp[0][5:], "/")
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

		sftSFTPData := models.DirListing{
			Name:          sftpPath,
			ChildrenDirs:  childrenDirs,
			ChildrenFiles: childrenFiles,
		}
		this.Data["fileData"] = sftSFTPData
	}
	return nil
}
