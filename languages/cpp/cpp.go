package cpp

import "github.com/palcode-oss/container-init/runner"

func Run() {
	params := runner.RunParams{
		Command:        "g++ main.cpp && ./main.cpp",
		ModuleFunction: runner.NoopModuleFunction,
	}
	runner.Run(params)
}
