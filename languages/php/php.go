package php

import "github.com/palcode-oss/container-init/runner"

func Run() {
	params := runner.RunParams{
		Command:        "php index.php",
		ModuleFunction: runner.NoopModuleFunction,
	}
	runner.Run(params)
}
