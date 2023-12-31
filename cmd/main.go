package main

import (
	"log"
	"os"

	"github.com/radenrishwan/gothon"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		gothon.Prompt()
	} else if len(args) == 2 {
		if err := gothon.Run(args[1]); err != nil {
			log.Fatalln(err)
		}
	} else {
		gothon.Help()
	}
}
