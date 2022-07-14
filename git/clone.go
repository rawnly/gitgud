package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
)

type CloneOptions struct {
	Local       bool   `json:"local" flag:"l"`
	Shared      bool   `json:"shared" flag:"s"`
	NoHardlinks bool   `json:"no-hardlinks" flag:"no-hardlinks"`
	Progress    bool   `json:"progress" flag:"progress"`
	NoCheckout  bool   `json:"no-checkout" flag:"n"`
	Bare        bool   `json:"bare" flag:"bare"`
	Mirror      bool   `json:"mirror" flag:"mirror"`
	Recursive   bool   `json:"recursive" flag:"recursive"`
	Branch      string `json:"branch" flag:"-b"`
	Quiet       bool   `json:"quiet" flag:"q"`
	Verbose     bool   `json:"verbose" flag:"v"`
	Directory   string
}

// Clone `git clone`
func Clone(repository string, options *CloneOptions) run.Runnable {
	opts := util.GetOptions(&options)
	var args []string

	for _, v := range util.StringifyOptions(opts) {
		args = append(args, v)
	}

	args = append(args, repository)

	// util.GetOptions ignores fields that don't have a flag tag
	if options.Directory != "" {
		args = append(args, options.Directory)
	}

	return run.Git("clone", args...)
}
