package prolog

import "github.com/palcode-oss/runner-common/runner"

func Run() {
	params := runner.RunParams{
		Command:        "swipl main.pl",
		ModuleFunction: runner.NoopModuleFunction,
	}

	runner.Run(params)
}
