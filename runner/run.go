package runner

import (
	"fmt"
	"os"
	"os/exec"
)

type ModuleFunction func(projectPath string) ([]string, bool, bool)
type EnvSetupFunction func(projectPath string)
type ModuleCommandGenerator func(modules []string) string

func Run(
	command string,
	moduleFunction ModuleFunction,
	moduleCommandGenerator ModuleCommandGenerator,
	envSetupFunction EnvSetupFunction,
) {
	const projectPath = "/usr/src/app"
	timeoutString := os.Args[1]

	modules, canRemoveEnv, success := moduleFunction(projectPath)
	if len(modules) == 0 {
		if canRemoveEnv {
			removeEnv(projectPath)
		}

		runCode(command, timeoutString, projectPath)
		os.Exit(0)
		return
	}

	if !success {
		fmt.Println("Warning: module detection crashed. No modules could be installed.")
		fmt.Println("This probably indicates a syntax error on your part.")
		fmt.Println("I'll run your code now so you can see what happened:")

		runCommand := exec.Command(
			"bash",
			"-c",
			"timeout 3s "+command,
		)

		runCommand.Stdout = os.Stdout
		runCommand.Stderr = os.Stderr
		_ = runCommand.Run()
		os.Exit(1)
		return
	}

	envSetupFunction(projectPath)
	moduleCommand := moduleCommandGenerator(modules)
	runInstall(moduleCommand, projectPath)
	runCode(command, timeoutString, projectPath)
	os.Exit(0)
}
