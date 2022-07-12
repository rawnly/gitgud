# gitgud
Git library with DX in mind

## Installation
```shell
go get -u github.com/rawnly/gitgud
```

## Usage
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

## Available Commands
#### Start a working area
- [ ] Clone 
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
