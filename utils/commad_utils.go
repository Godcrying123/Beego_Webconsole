package utils

import "os/exec"

func Command_Exec() error {
	command := exec.Command("/bin/bash", "-c", "systemctl staus firewalld")
	_, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	if err := command.Start(); err != nil {

	}

	return nil
}
