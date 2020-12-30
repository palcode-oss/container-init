package _go

import (
	"errors"
	"github.com/palcode-oss/container-init/runner"
	"os"
	"os/exec"
	"path"
)

func Run() {
	params := runner.RunParams{
		Command: "go run main.go",
		ModuleFunction: func(_ string) ([]string, bool, bool) {
			return []string{"keep"}, false, true
		},
		ModuleCommandGenerator: func(_ []string) string {
			return "go get -d ."
		},
		EnvSetupFunction: func(projectPath string) {
			_, err := os.Stat(
				path.Join(projectPath, "go.mod"),
			)

			if errors.Is(err, os.ErrNotExist) {
				modCommand := exec.Command(
					"go",
					"mod",
					"init",
					"app/v2",
				)
				modCommand.Stdout = os.Stdout
				modCommand.Dir = projectPath
				_ = modCommand.Run()
			}
		},
	}

	runner.Run(params)
}
