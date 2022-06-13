package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		gpkg.action.usage()
		os.Exit(1)
	}
}
