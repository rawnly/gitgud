package run

import (
	"os/exec"
	"testing"
)

var cmd = CmdWithStderr{exec.Command("gitgud")}

func TestCmdWithStderr_Run(t *testing.T) {
	// it should fail
	if err := cmd.Run(); err != nil {
		return
	}

	t.Fail()
}
