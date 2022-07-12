package run

import "os/exec"

func Git(command string, args ...string) (string, error) {
	arguments := []string{command}

	for _, v := range args {
		arguments = append(arguments, v)
	}

	cmd := exec.Command("git", arguments...)
	c := PrepareCmd(cmd)

	out, err := c.Output()

	return string(out), err
}
