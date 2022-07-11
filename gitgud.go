package gitgud

import (
	"github.com/rawnly/gitgud/util"
	"log"
	"os/exec"
	"strings"
)

func Git(command string, args ...string) (string, error) {
	arguments := []string{command}

	for _, v := range args {
		arguments = append(arguments, v)
	}

	log.Printf("Executing: git %s\n", strings.Join(arguments, " "))
	out, err := exec.Command("git", arguments...).Output()

	if err != nil {
		return "", err
	}

	return string(out), err
}

// SetConfig `git config <key>`
func SetConfig(key string, value string) (string, error) {
	output, err := exec.Command("git", "config", key, value).Output()

	if err != nil {
		return "", err
	}

	return string(output), err
}

// Config `git config <key> <value>`
func Config(key string) (string, error) {
	out, err := exec.Command("git", "config", key).Output()

	if err != nil {
		return "", err
	}

	return string(out), err
}

type CommitOptions struct {
	Message string `json:"message" flag:"m"`
	All     bool   `json:"all" flag:"a"`
}

// Commit `git commit -n -a -m [message]`
func Commit(options CommitOptions) (string, error) {
	opts := util.GetOptions(options)
	args := util.StringifyOptions(opts)

	return Git("commit", args...)
}

type PushOptions struct {
	Force       bool   `json:"force" flag:"f"`
	Delete      bool   `json:"delete" flag:"d"`
	Quiet       bool   `json:"quiet" flag:"q"`
	Verbose     bool   `json:"verbose" flag:"v"`
	SetUpstream string `json:"set_upstream" flag:"u"`
}

// Push `git push origin ${branch}`
func Push(branch string, options *PushOptions) (string, error) {
	opts := util.GetOptions(&options)
	args := []string{
		"origin", branch,
	}

	for _, v := range util.StringifyOptions(opts) {
		args = append(args, v)
	}

	return Git("commit", args...)
}

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

type StatusOptions struct {
	Short     bool   `flag:"s"`
	Branch    string `flag:"b"`
	ShowStash bool   `flag:"show-stash"`
	Long      bool   `flag:"long"`
	Verbose   bool   `flag:"v"`
}

// Status `git status -s`
func Status(options StatusOptions) ([]string, error) {
	colorUiConfig, err := Config("color.ui")

	if err != nil {
		if _, err := SetConfig("color.ui", "auto"); err != nil {
			return nil, err
		}

		return nil, err
	}

	if _, err := SetConfig("color.ui", "always"); err != nil {
		return nil, err
	}

	args := util.StringifyOptions(util.GetOptions(options))
	b, err := Git("status", args...)

	if err != nil {
		return nil, err
	}

	if _, err := SetConfig("color.ui", strings.TrimSpace(strings.Trim(colorUiConfig, "\n"))); err != nil {
		return nil, err
	}

	output := util.RemoveEmptyStrings(strings.Split(string(b), "\n"))

	return output, nil
}

// Diff `git diff`
func Diff() (string, error) {
	output, err := exec.Command("git", "diff").Output()

	if err != nil {
		return "", err
	}

	return string(output), err
}

// AddAll `git add -A .`
func AddAll() (string, error) {
	output, err := exec.Command("git", "add", "-A", ".").Output()

	if err != nil {
		return "", err
	}

	return string(output), err
}
