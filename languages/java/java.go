package java

import "github.com/palcode-oss/runner-common/runner"

func Run() {
	params := runner.RunParams{
		Command:        "java Main.java",
		ModuleFunction: runner.NoopModuleFunction,
	}

	runner.Run(params)
}
