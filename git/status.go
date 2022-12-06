package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
)

type StatusOptions struct {
	Short     bool   `flag:"s"`
	Branch    string `flag:"b"`
	ShowStash bool   `flag:"show-stash"`
	Long      bool   `flag:"long"`
	Verbose   bool   `flag:"v"`
	Path      string
}

// Status `git status -s .`
func Status(options *StatusOptions) run.Runnable {
	var args []string
	if options != nil {
		args = util.StringifyOptions(util.GetOptions(*options))
	}

	pathspec := "."
	if options.Path != "" {
		pathspec = options.Path
	}

	args = append(args, pathspec)

	return run.Git("status", args...)
}
