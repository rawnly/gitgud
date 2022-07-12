package run

import "os/exec"

func Git(command string, args ...string) Runnable {
	arguments := []string{command}

	for _, v := range args {
		arguments = append(arguments, v)
	}

	cmd := exec.Command("git", arguments...)

	return PrepareCmd(cmd)
}
