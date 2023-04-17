package git

import (
	"github.com/rawnly/gitgud/run"
	"github.com/rawnly/gitgud/util"
)

type CheckoutOptions struct {
	Create  bool `flag:"b"`
	Reset   bool `flag:"B"`
	Detach  bool `json:"detach" flag:"d"`
	Force   bool `json:"force" flag:"f"`
	Ours    bool `json:"ours" flag:"2"`
	Theirs  bool `json:"theirs" flag:"3"`
	Patch   bool `json:"patch" flag:"p"`
	Quiet   bool `json:"quiet" flag:"q"`
	Guess   bool `json:"guess"`
	Overlay bool `json:"overlay"`
	Reflog  bool `flag:"l"`
}

func Checkout(branch string, options CheckoutOptions) run.Runnable {
	opts := util.GetOptions(&options)
	args := util.StringifyOptions(opts)

	args = append(args, branch)

	return run.Git("checkout", args...)
}
