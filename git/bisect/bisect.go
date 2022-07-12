package bisect

import "github.com/rawnly/gitgud/run"

func Start() (string, error) {
	return run.Git("bisect", "start")
}

func Good(commit string) (string, error) {
	return run.Git("bisect", "good", commit)
}

func Bad(commit string) (string, error) {
	return run.Git("bisect", "bad", commit)
}

func Reset(commit string) (string, error) {
	return run.Git("bisect", "reset", commit)
}

func Run(command string) (string, error) {
	return run.Git("bisect", "run", "command")
}

func Skip(commit string) (string, error) {
	return run.Git("bisect", "skip", commit)
}

type Runner struct {
	Start func() (string, error)
	Bad   func(commit string) (string, error)
	Good  func(commit string) (string, error)
	Skip  func(commit string) (string, error)
	Reset func(commit string) (string, error)
	Run   func(command string) (string, error)
}

var R = Runner{
	Start: Start,
	Run:   Run,
	Bad:   Bad,
	Good:  Good,
	Reset: Reset,
	Skip:  Skip,
}
