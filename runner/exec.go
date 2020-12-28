package runner

import (
	"os"
	"os/exec"
	"strings"
)

type runCommandParams struct {
	command, timeout, directory string
	allowInput                  bool
}

func runCommand(params runCommandParams) {
	command := strings.ReplaceAll(params.command, `"`, `\"`)
	if params.timeout != "" {
		command = "timeout " + params.timeout + " " + command
	}

	instance := exec.Command(
		"bash",
		"-c",
		`"`+command+`"`,
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
