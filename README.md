# Gitgud
A tiny library to run git commands.

Sometimes working with `exec.Cmd` might not be easy as expected. Here gitgud comes to help you.
Every command has a standard `Runnable` output. It is similar to the `exec.Cmd` interface but with some tweaking.
If you need more customisation you can always use the original `exec.Cmd` instance and tweak it yourself.

## Features
- Ease of use
- Handle command errors
- Supports custom `stderr`, `stdout` and `stdin`
- Run directly in the terminal with `RunInTerminal`


## Usage
```shell
go get -u github.com/rawnly/gitgud
```
```go
package main

import (
	"fmt"
	"github.com/rawnly/gitgud/git"
)

func main() {
	// equivalent to "git status -s" 
	status, err := git.Status(&git.StatusOptions{
		Short: true,
	}).Output()

	if err != nil {
		panic(err.Error())
	}
	
	fmt.Println(status)
} 
```

Advanced Example
```go
package main 

import "github.com/rawnly/gitgud/git"
import "os"

func main() {
	diffCmd := git.Commit("docs(readme): updated example", nil)
	
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	
	diffCmd.Cmd.Stdout = stdout
	diffCmd.Cmd.Stderr = stderr
	
	if err := diffCmd.Run(); err != nil {
		panic(err)
	}
}
```

## Available Commands
#### Start a working area
- [x] Clone 
- [ ] Init

#### Work on the current change
- [ ] Add
- [ ] Mv
- [ ] Restore
- [ ] Rm

#### Examine the history and state
- [x] Bisect
- [ ] Diff
- [ ] Grep
- [ ] Log
- [ ] Show 
- [x] Status

#### Grow, mark and tweak your common history
- [ ] Branch
- [x] Commit
- [ ] Merge
- [ ] Rebase
- [ ] Reset
- [ ] Switch
- [ ] Tag

#### Collaborate
- [ ] Fetch
- [x] Push
- [ ] Pull

#### Others
- [x] Config, SetConfig
