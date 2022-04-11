package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hasahmad/go-api/internal/api"
)

var (
	version   string
	buildTime string
)

func main() {
	displayVersion := flag.Bool("version", false, "Display version and exit")
	flag.Parse()

	// If the version flag value is true, then print out the version number and
	// immediately exit.
	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	api.StartServer()
}
