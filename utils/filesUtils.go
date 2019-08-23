package utils

import (
	"archive/zip"
	"bufio"
	"fmt"
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

func CompressFiles(baseFolder, zipFilePath string) (err error) {
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
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
	return nil
}
