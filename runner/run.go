package runner

import (
	"fmt"
	"os"
)

type ModuleFunction func(projectPath string) ([]string, bool, bool)
type EnvSetupFunction func(projectPath string)
type ModuleCommandGenerator func(modules []string) string

type RunParams struct {
	Command                string
	ModuleFunction         ModuleFunction
	ModuleCommandGenerator ModuleCommandGenerator
	EnvSetupFunction       EnvSetupFunction
}

func Run(params RunParams) {
	const projectPath = "/usr/src/app"
	timeoutString := os.Args[1]

	modules, canRemoveEnv, success := params.ModuleFunction(projectPath)
	if len(modules) == 0 {
		if canRemoveEnv {
			removeEnv(projectPath)
		}

		ClearScreen()
		runCommand(runCommandParams{
			command:    params.Command,
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
			command:    params.Command,
			timeout:    "3s",
			directory:  projectPath,
			allowInput: false,
		})
		os.Exit(1)
		return
	}

	params.EnvSetupFunction(projectPath)
	moduleCommand := params.ModuleCommandGenerator(modules)
	runCommand(runCommandParams{
		command:    moduleCommand,
		directory:  projectPath,
		allowInput: false,
	})

	ClearScreen()
	runCommand(runCommandParams{
		command:    params.Command,
		timeout:    timeoutString,
		directory:  projectPath,
		allowInput: true,
	})
	os.Exit(0)
}
