package utils

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/astaxie/beego"
)

func init() {}

func CommandExec(command string) (result string, err error) {
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
				return builder.String(), nil
			}
		}
		builder.Write(output)
		builder.Write([]byte("\n"))
	}
}
