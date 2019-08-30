package utils

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
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

func compressFiles(baseFolder, zipFilePath string) (err error) {
	// Get a Buffer to Write To
	outFile, err := os.Create(zipFilePath)
	if err != nil {
		beego.Error(err)
		return err
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	err = addFiles(w, baseFolder, "")

	if err != nil {
		beego.Error(err)
		return err
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		beego.Error(err)
		return err
	}
	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) (err error) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		beego.Error(err)
		return err
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				beego.Error(err)
				return err
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				beego.Error(err)
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				beego.Error(err)
				return err
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			// fmt.Println("Recursing and Adding SubDir: " + file.Name())
			// fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
	return nil
}

func FileDownLoad(fileName string, raw io.Reader) (err error) {
	reader := bufio.NewReaderSize(raw, 1024*32)
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)

	buff := make([]byte, 32*1024)
	written := 0
	go func() {
		for {
			nr, er := reader.Read(buff)
			if nr > 0 {
				// beego.Info(buff)
				nw, ew := writer.Write(buff[0:nr])
				if nw > 0 {
					written += nw
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortWrite
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		if err != nil {
			panic(err)
		}
	}()

	spaceTime := time.Second * 1
	ticker := time.NewTicker(spaceTime)
	lastWtn := 0
	stop := false

	for {
		select {
		case <-ticker.C:
			speed := written - lastWtn
			log.Printf("[*] Speed %s / %s \n", bytesToSize(speed), spaceTime.String())
			if written-lastWtn == 0 {
				ticker.Stop()
				stop = true
				break
			}
			lastWtn = written
		}
		if stop {
			break
		}
	}
	return nil
}

func bytesToSize(length int) string {
	var k = 1024 // or 1024
	var sizes = []string{"Bytes", "KB", "MB", "GB", "TB"}
	if length == 0 {
		return "0 Bytes"
	}
	i := math.Floor(math.Log(float64(length)) / math.Log(float64(k)))
	r := float64(length) / math.Pow(float64(k), i)
	return strconv.FormatFloat(r, 'f', 3, 64) + " " + sizes[int(i)]
}

func DirAndFileDownLoad(writer http.ResponseWriter, request *http.Request, fileDir, fileName string) (err error) {
	file, err := os.Open(fileDir + fileName)

	if err != nil {
		return err
	}
	defer file.Close()
	statinfo, err := file.Stat()
	if err != nil {
		return err
	}
	if statinfo.IsDir() {
		zipFileName := strings.Trim(fileName, "/") + ".zip"
		compressFiles(fileDir+fileName, fileDir+zipFileName)
		err = downloadFile(writer, request, fileDir, zipFileName)
		if err != nil {
			return err
		}
		err = os.Remove(fileDir + zipFileName)
		if err != nil {
			beego.Error(err)
			return err
		}
	} else {
		err = downloadFile(writer, request, fileDir, fileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadFile(writer http.ResponseWriter, request *http.Request, filePath, fileName string) (err error) {
	beego.Info("[*] Filename " + fileName)

	//Check if file exists and open
	Openfile, err := os.Open(filePath + fileName)
	defer Openfile.Close() //Close after function return
	// writer := this.Ctx.ResponseWriter.ResponseWriter
	if err != nil {
		//File not found, send 404
		http.Error(writer, "File not found.", 404)
		return err
	}

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	Openfile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Send the headers
	writer.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	writer.Header().Set("Content-Type", FileContentType)
	writer.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	Openfile.Seek(0, 0)
	io.Copy(writer, Openfile) //'Copy' the file to the client
	return nil
}
