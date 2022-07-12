package git

import "github.com/rawnly/gitgud/run"

// SetConfig `git config <key>`
func SetConfig(key string, value string) (string, error) {
	return run.Git("config", key, value)
}

// Config `git config <key> <value>`
func Config(key string) (string, error) {
	return run.Git("config", key)
}
