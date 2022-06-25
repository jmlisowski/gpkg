package main

import (
	"os"

	"github.com/jmlisowski/gpkg/action"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		action.Usage()
		os.Exit(1)
	}
	if len(args) > 0 {
		if args[0] == "init" {
			action.Init()
		} else if args[0] == "install" {
			if len(args) < 2 {
				println("must specify a package!")
			} else if len(args) == 2 {
				action.Install(args[1])
			} else {
				println("only install one package at once!")
			}
		}
	}
}
