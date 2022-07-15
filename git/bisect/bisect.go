package bisect

import "github.com/rawnly/gitgud/run"

func Start() run.Runnable {
	return run.Git("bisect", "start")
}

func Good(commit string) run.Runnable {
	return run.Git("bisect", "good", commit)
}

func Bad(commit string) run.Runnable {
	return run.Git("bisect", "bad", commit)
}

func Reset(commit string) run.Runnable {
	return run.Git("bisect", "reset", commit)
}

func Run(command string) run.Runnable {
	return run.Git("bisect", "run", "command")
}

func Skip(commit string) run.Runnable {
	return run.Git("bisect", "skip", commit)
}

type Runner struct {
	Start func() run.Runnable
	Bad   func(commit string) run.Runnable
	Good  func(commit string) run.Runnable
	Skip  func(commit string) run.Runnable
	Reset func(commit string) run.Runnable
	Run   func(command string) run.Runnable
}
