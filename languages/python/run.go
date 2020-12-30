package python

import (
	"fmt"
	"github.com/palcode-oss/container-init/languages/python/detect_modules"
	"github.com/palcode-oss/container-init/runner"
	"os"
	"os/exec"
	"path"
)

func Run() {
	params := runner.RunParams{
		Command: "source env/bin/activate 2>/dev/null ; python main.py",
		ModuleFunction: func(projectPath string) ([]string, bool, bool) {
			return detect_modules.DetectModules(projectPath)
		},
		ModuleCommandGenerator: func(modules []string) string {
			moduleCommand := "source env/bin/activate ; python -m pip install "
			for _, module := range modules {
				moduleCommand += module + " "
			}
			return moduleCommand
		},
		EnvSetupFunction: func(projectPath string) {
			envActivatePath := path.Join(projectPath, "env", "bin", "activate")

			if _, err := os.Stat(envActivatePath); os.IsNotExist(err) {
				fmt.Println("Setting up environment...")
				envPath := path.Join(projectPath, "env")
				venvCommand := exec.Command(
					"python",
					"-m",
					"venv",
					envPath,
				)
				venvCommand.Stdout = os.Stdout
				_ = venvCommand.Run()
			}
		},
	}

	runner.Run(params)
}
