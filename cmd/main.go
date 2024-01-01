package main

import (
	"flag"
	"log"
	"os"

	"github.com/radenrishwan/gothon"
)

var (
	version = flag.String("version", "0.0.1", "gothon version")
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
