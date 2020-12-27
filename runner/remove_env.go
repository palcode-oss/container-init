package runner

import (
	"os"
	"path"
)

func removeEnv(directory string) {
	envRelatedStrings := [...]string{
		"env/",
		"node_modules/",
		"requirements.txt",
		"package.json",
		"yarn.lock",
	}

	for _, pathName := range envRelatedStrings {
		_ = os.RemoveAll(
			path.Join(directory, pathName),
		)
	}
}
