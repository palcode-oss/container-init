package runner

import (
	"os"
	"os/exec"
)

func runInstall(moduleCommand, directory string) {
	installCommand := exec.Command(
		"bash",
		"-c",
		"cd "+directory+" ; "+moduleCommand,
	)
	installCommand.Stdout = os.Stdout
	installCommand.Stderr = os.Stderr
	_ = installCommand.Run()
}
