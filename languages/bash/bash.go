package bash

import "github.com/palcode-oss/runner-common/runner"

func Run() {
	params := runner.RunParams{
		Command:        "bash main.sh",
		ModuleFunction: runner.NoopModuleFunction,
	}

	runner.Run(params)
}
