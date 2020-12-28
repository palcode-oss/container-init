package runner

import "os/exec"

func withEnv(command, directory string) *runCommand.Cmd {
	return exec.Command(
		"bash",
		"-c",
		"cd "+directory+" ; source env/bin/activate ; "+command,
	)
}
