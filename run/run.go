package run

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

//Custom implementation of the cmd struct
//Inspired by github.com/cli/cli/internal/run

type Runnable interface {
	Output() ([]byte, error)
}

var PrepareCmd = func(cmd *exec.Cmd) Runnable {
	return &CmdWithStderr{cmd}
}

type CmdError struct {
	Stderr *bytes.Buffer
	Args   []string
	Err    error
}

func (e CmdError) Error() string {
	msg := e.Stderr.String()

	if msg != "" && !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	return fmt.Sprintf("%s%s: %s", msg, e.Args[0], e.Err)
}

type CmdWithStderr struct {
	*exec.Cmd
}

func (c CmdWithStderr) Output() ([]byte, error) {
	if c.Cmd.Stderr != nil {
		return c.Cmd.Output()
	}

	errStream := &bytes.Buffer{}
	c.Cmd.Stderr = errStream

	out, err := c.Cmd.Output()

	if err != nil {
		err = &CmdError{errStream, c.Cmd.Args, err}
	}

	return out, err
}
