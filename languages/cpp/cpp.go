package cpp

import "github.com/palcode-oss/container-init/runner"

func Run() {
	params := runner.RunParams{
		Command:        "g++ main.cpp && ./a.out",
		ModuleFunction: runner.NoopModuleFunction,
	}
	runner.Run(params)
}
