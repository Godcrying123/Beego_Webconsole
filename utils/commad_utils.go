package utils

import "os/exec"

func init() {}

func Command_Exec() error {
	command := exec.Command("/bin/bash", "-c", "df -ah")
	_, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	if err := command.Start(); err != nil {

	}

	return nil
}
