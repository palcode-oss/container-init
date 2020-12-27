package runner

import "os/exec"

func withEnv(command, directory string) *exec.Cmd {
	return exec.Command(
		"bash",
		"-c",
		"cd "+directory+" ; source env/bin/activate ; "+command,
	)
}
