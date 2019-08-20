package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/astaxie/beego"
)

func CommandExecReturnString1() {
	cmd := exec.Command("/bin/bash", "-c", `df -lh`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	beego.Info(fmt.Sprintf("stdout %s", bytes))
	//fmt.Printf("stdout:\n\n %s", bytes)
}

func CommandExecReturnString(command string) (result string, err error) {
	var builder strings.Builder
	commandexec := exec.Command("/bin/bash", "-c", command)
	stdout, err := commandexec.StdoutPipe()
	if err != nil {
		beego.Error(err)
		return "", err
	}

	if err := commandexec.Start(); err != nil {
		beego.Error(err)
		return "", err
	}
	outputBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				beego.Error(err)
				return "", err
			} else {
				return strings.Trim(builder.String(), "\n"), nil
			}
		}
		builder.Write(output)
		builder.Write([]byte("\n"))
	}
}

func CommandExecReturnSlices(command string) (resultslice []string, err error) {
	commandexec := exec.Command("/bin/bash", "-c", command)
	stdout, err := commandexec.StdoutPipe()
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	if err := commandexec.Start(); err != nil {
		beego.Error(err)
		return nil, err
	}
	outputBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				beego.Error(err)
				return nil, err
			} else {
				return resultslice, nil
			}
		}
		resultslice = append(resultslice, fmt.Sprintf("%s", output))
		//builder.Write([]byte("\n"))
	}
}

func CommandExecWithoutPipe(comand string) {
	in := bytes.NewBuffer(nil)
	cmd := exec.Command("sh")
	cmd.Stdin = in
	go func() {
		in.WriteString("systemctl list-units -allgrep 'loaded units listed'> test.txt\n")
		in.WriteString("exit\n")
	}()
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
