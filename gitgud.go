package gitgud

import (
	"os/exec"
	"strings"
)

// CurrentBranch Get current branch
func CurrentBranch() string {
	var currentBranch string

	b, err := exec.Command("git", "branch").Output()

	if err != nil {
		return currentBranch
	}

	branch := string(b)

	output := strings.ReplaceAll(branch, " ", "")

	for _, b := range strings.Split(output, "\n") {
		if strings.Contains(b, "*") {
			currentBranch = b
		}
	}

	return strings.ReplaceAll(currentBranch, "*", "")
}
