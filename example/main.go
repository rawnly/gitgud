package main

import (
	"fmt"
	"github.com/rawnly/gitgud/git"
)

func main() {
	out, err := git.Bisect.Start()
	//out, err := git.Status(&git.StatusOptions{
	//	Short: true,
	//})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(out)
}
