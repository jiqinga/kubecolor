package main

import (
	"errors"
	"os"

	"github.com/jiqinga/kubecolor/command"
)

// this is overridden on build time by GoReleaser
var Version = "unset"

func main() {
	env := os.Environ()
	err := command.Run(os.Args[1:], Version, env)
	if err != nil {
		var ke *command.KubectlError
		if errors.As(err, &ke) {
			os.Exit(ke.ExitCode)
		}
		os.Exit(1)
	}
}
