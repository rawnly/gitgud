package run

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Custom implementation of the cmd struct
//Inspired by github.com/cli/cli/internal/run

type Runnable interface {
	// Output Returns output and error
	Output() ([]byte, error)
	// Run Ignores the output
	Run() error
	// RunInTerminal uses `os.stdin`, `os.stdout`, `os.stderr` as input/output.
	RunInTerminal() error
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

func (c CmdWithStderr) Run() error {
	if c.Cmd.Stderr != nil {
		return c.Cmd.Run()
	}

	errStream := &bytes.Buffer{}
	c.Cmd.Stderr = errStream

	err := c.Cmd.Run()

	if err != nil {
		err = &CmdError{errStream, c.Cmd.Args, err}
	}

	return err
}

func (c CmdWithStderr) RunInTerminal() error {
	c.Cmd.Stdout = os.Stdout
	c.Cmd.Stderr = os.Stderr
	c.Cmd.Stdin = os.Stdin

	return c.Cmd.Run()
}