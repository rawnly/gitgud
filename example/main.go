package main

import (
	"github.com/rawnly/gitgud/run"
)

func main() {
	commit := run.NewGitBuilder("clone").
		Arg("--local").
		Arg("git@github.com:rawnly/gitgud").
		BoolFlag("--recursive", true).
		Arg("local-folder").
		Build()

	if err := commit.Run(); err != nil {
		panic(err.Error())
	}
}
