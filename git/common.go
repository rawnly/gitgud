package git

import "github.com/rawnly/gitgud/git/bisect"

func Bisect() bisect.Runner {
	return bisect.Runner{
		Start: bisect.Start,
		Run:   bisect.Run,
		Bad:   bisect.Bad,
		Good:  bisect.Good,
		Reset: bisect.Reset,
		Skip:  bisect.Skip,
	}
}
