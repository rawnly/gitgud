package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
)

type PushOptions struct {
	Force       bool   `json:"force" flag:"f"`
	Delete      bool   `json:"delete" flag:"d"`
	Quiet       bool   `json:"quiet" flag:"q"`
	Verbose     bool   `json:"verbose" flag:"v"`
	SetUpstream string `json:"set_upstream" flag:"u"`
}

// Push `git push origin ${branch}`
func Push(branch string, options *PushOptions) run.Runnable {
	opts := util.GetOptions(&options)
	args := []string{
		"origin", branch,
	}

	for _, v := range util.StringifyOptions(opts) {
		args = append(args, v)
	}

	return run.Git("commit", args...)
}
