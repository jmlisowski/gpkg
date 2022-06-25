package action

import (
	"log"
	"os"
	"os/exec"

	"github.com/jmlisowski/gpkg/dfile"
)

//Usage function tells the user how to use gpkg
func Usage() {
	println("Usage:")
	println("  gpkg <command> [<args>]")
	println("")
	println("Commands:")
	println("  init      Sets up the package manager(RUN THIS THE FIRST USE)")
	println("  install   Installs a package")
	println("  remove    Removes a package")
	println("  run       Runs a package")
}

//Init function makes gpkg ready to use
func Init() {
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

//Install function installs a package from github
func Install(pkg string) {
	// make a variable for the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	//download a file from the jmlisowski/gpkg repository
	//and install it
	dfile.DownloadFile("https://raw.githubusercontent.com/jmlisowski/gpkg/packages/"+pkg+".tar", homeDir+"/.gpkg/"+pkg+".tar")
	cmd := exec.Command("tar", "-xvf", homeDir+"/.gpkg/"+pkg+".tar", "-C", homeDir+"/.gpkg/")
	cmderr := cmd.Run()
	if cmderr != nil {
		log.Fatal(err)
	}

	path := homeDir + "/.gpkg/" + pkg + ".tar"
	rmerr := os.Remove(path)
	if rmerr != nil {
		log.Fatal(err)
	}
}

//remove function
func Remove(pkg string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	rmerr := os.RemoveAll(homeDir + "/.gpkg/" + pkg + "/")

	if rmerr != nil {
		panic(rmerr)
	}
}
