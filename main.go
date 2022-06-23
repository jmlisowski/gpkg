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
			action.Install()
		}
	}
}
