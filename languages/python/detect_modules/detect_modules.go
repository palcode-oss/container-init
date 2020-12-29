package detect_modules

import (
	"github.com/go-python/gpython/ast"
	"github.com/go-python/gpython/parser"
	"github.com/thoas/go-funk"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func checkIsLocalFile(module string, files []fs.FileInfo) bool {
	isLocalFile := false
	for _, file := range files {
		if file.Name() == module+".py" {
			isLocalFile = true
		}
	}
	return isLocalFile
}

func checkIsBuiltIn(module string) bool {
	isBuiltIn := false
	for _, builtin := range PythonBuiltIns {
		if module == builtin {
			isBuiltIn = true
		}
	}
	return isBuiltIn
}

func checkIsAlreadyInstalled(module string) bool {
	_, fileErr := os.Stat(
		path.Join("/usr/src/app/env/lib/python3.9/site-packages", module+".py"),
	)

	_, directoryErr := os.Stat(
		path.Join("/usr/src/app/env/lib/python3.9/site-packages", module),
	)

	if os.IsNotExist(fileErr) && os.IsNotExist(directoryErr) {
		return false
	} else {
		return true
	}
}

func removeUselessModules(modules []string, files []fs.FileInfo) ([]string, bool) {
	var usefulModules []string
	canRemoveEnv := true
	for _, module := range modules {
		isLocalFile := checkIsLocalFile(module, files)
		if isLocalFile {
			continue
		}

		isBuiltIn := checkIsBuiltIn(module)
		if isBuiltIn {
			continue
		}

		isAlreadyInstalled := checkIsAlreadyInstalled(module)
		if isAlreadyInstalled {
			canRemoveEnv = false
			continue
		}

		usefulModules = append(usefulModules, module)
	}

	return usefulModules, canRemoveEnv
}

func DetectModules(directory string) ([]string, bool, bool) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatal("Couldn't read directory")
	}

	var modules []string
	for _, file := range files {
		filePath := path.Join(directory, file.Name())

		isPythonFile := strings.HasSuffix(filePath, ".py")
		if !isPythonFile {
			continue
		}

		fileContentsBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}

		fileContents := string(fileContentsBytes)
		parsedFile, err := parser.ParseString(fileContents, "exec")
		if err != nil {
			return []string{}, false, false
		}

		ast.Walk(parsedFile, func(rootNode ast.Ast) bool {
			switch node := rootNode.(type) {
			case *ast.Import:
				for _, module := range node.Names {
					modules = append(modules, string(module.Name))
				}
			case *ast.ImportFrom:
				modules = append(modules, string(node.Module))
			}

			return true
		})
	}

	uniqueModules := funk.Uniq(modules).([]string)
	usefulModules, canRemoveEnv := removeUselessModules(uniqueModules, files)

	return usefulModules, canRemoveEnv, true
}
