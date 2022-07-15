package git

import (
	"fmt"
	"github.com/rawnly/gitgud/run"
)

type GitCommand struct {
	name      string
	arguments []string
}

type CommandBuilder interface {
	StringFlag(flag string, value string) *CommandBuilder
	BoolFlag(flag string, value bool) *CommandBuilder

	Build() run.Runnable
}

type GitCommandBuilder struct {
	CommandBuilder
	command *GitCommand
}

func (b *GitCommandBuilder) BoolFlag(flag string, value bool) *GitCommandBuilder {
	b.command.arguments = append(b.command.arguments, flag)
	b.command.arguments = append(b.command.arguments, fmt.Sprint(value))

	return b
}

func (b *GitCommandBuilder) Arg(arg string) *GitCommandBuilder {
	b.command.arguments = append(b.command.arguments, arg)

	return b
}

func (b *GitCommandBuilder) StringFlag(flag string, value string) *GitCommandBuilder {
	b.command.arguments = append(b.command.arguments, flag)
	b.command.arguments = append(b.command.arguments, value)

	return b
}

func (b *GitCommandBuilder) Build() run.Runnable {
	return run.Git(b.command.name, b.command.arguments...)
}

func NewGitBuilder(name string) *GitCommandBuilder {
	return &GitCommandBuilder{
		command: &GitCommand{name, []string{}},
	}
}
