package main

import (
	"fmt"
	"github.com/rawnly/gitgud/git"
	"log"
)

func main() {
	status, err := git.Status(&git.StatusOptions{
		Short: true,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(status)
}
