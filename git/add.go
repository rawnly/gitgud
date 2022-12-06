package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
)

type AddOptions struct {
	All  bool   `flag:"a"`
	Path string `arg:""`
}

func Add(options *AddOptions) run.Runnable {
	var args []string

	if options != nil {
		args = util.StringifyOptions(util.GetOptions(options))
		args = append(args, util.GetArguments(options)...)
	}

	return run.Git("add", args...)
}
