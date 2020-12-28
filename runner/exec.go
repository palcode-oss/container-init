package runner

import (
	"os"
	"os/exec"
)

type runCommandParams struct {
	command, timeout, directory string
	allowInput                  bool
}

func runCommand(params runCommandParams) {
	command := params.command
	program := "bash"
	arguments := []string{
		"-c",
		command,
	}

	if params.timeout != "" {
		program = "timeout"
		arguments = []string{
			params.timeout,
			"bash",
			"-c",
			command,
		}
	}

	instance := exec.Command(
		program,
		arguments...,
	)

	if params.directory != "" {
		instance.Dir = params.directory
	}

	instance.Stdout = os.Stdout
	instance.Stderr = os.Stderr

	if params.allowInput {
		instance.Stdin = os.Stdin
	}

	_ = instance.Run()
}
