package bash

import "github.com/palcode-oss/container-init/runner"

func Run() {
	params := runner.RunParams{
		Command:        "bash main.sh",
		ModuleFunction: runner.NoopModuleFunction,
	}

	runner.Run(params)
}
