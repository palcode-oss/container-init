package java

import "github.com/palcode-oss/container-init/runner"

func Run() {
	params := runner.RunParams{
		Command:        "java Main.java",
		ModuleFunction: runner.NoopModuleFunction,
	}

	runner.Run(params)
}
