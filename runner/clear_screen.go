package runner

import (
	"os"
	"os/exec"
)

func ClearScreen() {
	clearCommand := exec.Command("clear")
	clearCommand.Stdout = os.Stdout
	_ = clearCommand.Run()
}
