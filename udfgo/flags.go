package main

import (
	"fmt"
	"os"

	goflags "github.com/jessevdk/go-flags"
)

var flags struct {
	Verbose bool `short:"v" long:"verbose" description:"enable verbose logs"`

	Build cmdBuild `command:"build"`
}

func setupFlags() {
	// setup parser with our flags and custom options
	flagsParser := goflags.NewParser(&flags, goflags.HelpFlag|goflags.PrintErrors)

	// parse command line arguments
	args, err := flagsParser.Parse()
	if err != nil {
		// assert the err to be a flags.Error
		flagError, ok := err.(*goflags.Error)
		if !ok {
			// not a flags error
			os.Exit(1)
		}
		if flagError.Type == goflags.ErrHelp {
			// exitcode 0 when user asked for help
			os.Exit(0)
		}
		if flagError.Type == goflags.ErrUnknownFlag {
			fmt.Println("run with --help to view available options")
		}
		os.Exit(1)
	}

	// error on left-over arguments
	if len(args) > 0 {
		fmt.Printf("unexpected arguments: %s\n", args)
		os.Exit(0)
	}
}

func verbosef(format string, args ...interface{}) {
	if flags.Verbose {
		fmt.Printf(format, args...)
	}
}
