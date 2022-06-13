package main

import "os"

func usage() {
	println("Usage:")
	println("  gpkg <command> [<args>]")
	println("")
	println("Commands:")
	println("  init      Sets up the package manager(RUN THIS THE FIRST USE)")
	println("  install   Installs a package")
	println("  remove    Removes a package")
	println("  run       Runs a package")
}

func init() {
	//make a directory at ~/.gpkg to store the packages
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if err := os.Mkdir(homeDir+"/.gpkg", 0755); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
	println("gpkg is ready to use")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}
}
