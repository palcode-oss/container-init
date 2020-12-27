package runner

import (
	"os"
	"os/exec"
)

func runCode(command, timeoutString, directory string) {
	clearScreen()

	commandString := "timeout " + timeoutString + " " + command
	runCommand := exec.Command(
		"bash",
		"-c",
		"cd "+directory+" ; "+commandString,
	)

	runCommand.Stdout = os.Stdout
	runCommand.Stderr = os.Stderr
	runCommand.Stdin = os.Stdin
	_ = runCommand.Run()
}
