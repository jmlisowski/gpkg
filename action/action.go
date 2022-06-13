package action

import (
	"gpkg/dfile"
	"os"
)

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

func install() {
	// make a variable for the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	//download a file from the jmlisowski/gpkg repository
	//and install it
	dfile.DownloadFile("https://github.com/jmlisowski/gpkg/blob/main/packages/test_server.jar", "~/.gpkg/test_server.jar")
}
