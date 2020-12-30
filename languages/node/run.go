package node

import (
	"errors"
	"github.com/palcode-oss/container-init/languages/node/detect_modules"
	"github.com/palcode-oss/container-init/runner"
	"os"
	"os/exec"
	"path"
)

func Run() {
	params := runner.RunParams{
		Command: "node index.js",
		ModuleFunction: func(directory string) ([]string, bool, bool) {
			return detect_modules.DetectModules(directory)
		},
		ModuleCommandGenerator: func(modules []string) string {
			moduleCommand := "yarn add "
			for _, module := range modules {
				moduleCommand += module + " "
			}
			return moduleCommand
		},
		EnvSetupFunction: func(directory string) {
			_, err := os.Stat(
				path.Join(directory, "package.json"),
			)

			if errors.Is(err, os.ErrNotExist) {
				command := exec.Command(
					"yarn",
					"init",
					"-y",
				)
				_ = command.Run()
			}
		},
	}

	runner.Run(params)
}
