package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
	"strings"
)

type StatusOptions struct {
	Short     bool   `flag:"s"`
	Branch    string `flag:"b"`
	ShowStash bool   `flag:"show-stash"`
	Long      bool   `flag:"long"`
	Verbose   bool   `flag:"v"`
}

// Status `git status -s`
func Status(options *StatusOptions) (string, error) {
	colorUiConfig, err := Config("color.ui")

	if err != nil {
		if _, err := SetConfig("color.ui", "auto"); err != nil {
			return "", err
		}

		return "", err
	}

	if _, err := SetConfig("color.ui", "always"); err != nil {
		return "", err
	}

	var args []string
	if options != nil {
		args = util.StringifyOptions(util.GetOptions(*options))
	}

	b, err := run.Git("status", args...)

	if err != nil {
		return "", err
	}

	if _, err := SetConfig("color.ui", strings.TrimSpace(strings.Trim(colorUiConfig, "\n"))); err != nil {
		return "", err
	}

	return b, nil
}
