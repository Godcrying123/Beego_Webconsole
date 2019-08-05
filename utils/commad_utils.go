package utils

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/astaxie/beego"
)

func init() {}

func Command_Exec(command string) (result string, err error) {
	var builder strings.Builder
	command_exec := exec.Command("/bin/bash", "-c", command)
	stdout, err := command_exec.StdoutPipe()
	if err != nil {
		beego.Error(err)
		return "", err
	}

	if err := command_exec.Start(); err != nil {
		beego.Error(err)
		return "", err
	}
	beego.Info(stdout)
	outputBuf := bufio.NewReader(stdout)
	beego.Info(outputBuf)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				beego.Error(err)
			}
			return "", err
		}
		builder.Write(output)
		builder.Write([]byte("\n"))
	}

	return builder.String(), nil
}
