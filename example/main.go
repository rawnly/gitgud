package main

import (
	"fmt"
	"github.com/rawnly/gitgud"
	"log"
)

func main() {
	status, err := gitgud.Status(gitgud.StatusOptions{
		Verbose: true,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(status)
}
