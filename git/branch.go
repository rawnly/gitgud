package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
)

type BranchOptions struct {
	Delete      bool   `json:"delete" flag:"d"`
	All         bool   `json:"all" flag:"a"`
	SetUpStream string `json:"set_upstream" flag:"u"`
	Force       bool   `json:"force" flag:"f"`
	List        bool   `json:"list" flag:"l"`
}

func Branch(options BranchOptions) run.Runnable {
	opts := util.GetOptions(&options)
	args := []string{}

	for _, v := range util.StringifyOptions(opts) {
		args = append(args, v)
	}

	return run.Git("branch", args...)
}
