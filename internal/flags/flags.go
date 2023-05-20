package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/wigit-gh/webapp/internal/logging"
)

func usage() {
	fmt.Print(`
This executable runs the WIG!T Web Application backend.

Usage:

wwapp [arguments]

Supported arguments:

`)
	flag.PrintDefaults()
	os.Exit(1)
}

// Parse sets up the flags for the build executable.
func Parse() string {
	// use our usage function to display usage message if any error occurs during parsing.
	flag.Usage = usage

	// set the expected flags and default value.
	env := flag.String("env", "dev", `Specifies the run environment. Value is either "dev" or "prod"`)

	// Parse all command line flags.
	flag.Parse()

	// Configure global logger with specified environment.
	logFile := logging.ConfigureLogger(*env)
	if *env == "prod" && logFile != nil {
		logging.SetGinLogToFile(logFile)
	}

	return *env
}
