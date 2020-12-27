package runner

import (
	"os"
	"os/exec"
)

func clearScreen() {
	clearCommand := exec.Command("clear")
	clearCommand.Stdout = os.Stdout
	_ = clearCommand.Run()
}
