package runner

import (
	"fmt"
	"os"
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

		clearScreen()
		runCommand(runCommandParams{
			command:    command,
			timeout:    timeoutString,
			directory:  projectPath,
			allowInput: true,
		})
		os.Exit(0)
		return
	}

	if !success {
		fmt.Println("Warning: module detection crashed. No modules could be installed.")
		fmt.Println("This probably indicates a syntax error on your part.")
		fmt.Println("I'll run your code now so you can see what happened:")

		runCommand(runCommandParams{
			command:    command,
			timeout:    "3s",
			directory:  projectPath,
			allowInput: false,
		})
		os.Exit(1)
		return
	}

	envSetupFunction(projectPath)
	moduleCommand := moduleCommandGenerator(modules)
	runCommand(runCommandParams{
		command:    moduleCommand,
		directory:  projectPath,
		allowInput: false,
	})

	clearScreen()
	runCommand(runCommandParams{
		command:    command,
		timeout:    timeoutString,
		directory:  projectPath,
		allowInput: true,
	})
	os.Exit(0)
}
