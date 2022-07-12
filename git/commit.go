package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
)

type CommitOptions struct {
	Message string `json:"message" flag:"m"`
	All     bool   `json:"all" flag:"a"`
}

// Commit `git commit -n -a -m [message]`
func Commit(options CommitOptions) (string, error) {
	opts := util.GetOptions(options)
	args := util.StringifyOptions(opts)

	return run.Git("commit", args...)
}
