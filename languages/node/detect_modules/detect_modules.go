package detect_modules

import (
	"errors"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/js"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func checkIsAlreadyInstalled(module, directory string) bool {
	_, err := os.Stat(
		path.Join(directory, "node_modules", module),
	)

	return !errors.Is(err, os.ErrNotExist)
}

func checkIsLocalFile(module, directory string) bool {
	if !strings.HasSuffix(module, ".js") {
		module = module + ".js"
	}

	_, err := os.Stat(
		path.Join(directory, module),
	)
	return !errors.Is(err, os.ErrNotExist)
}

func checkIsBuiltin(module string) bool {
	for _, builtin := range BuiltinModules {
		if module == builtin {
			return true
		}
	}

	return false
}

func parseExpression(expression js.IExpr, modules *[]string) {
	switch expression := expression.(type) {
	case *js.CallExpr:
		if len(expression.Args.List) == 0 || expression.X.String() != "require" {
			return
		}

		argument := expression.Args.List[0]
		switch argument := argument.(type) {
		case *js.LiteralExpr:
			moduleName := string(argument.Data)
			trimmedModuleName := moduleName[1 : len(moduleName)-1]
			*modules = append(*modules, trimmedModuleName)
		}
	}
}

func DetectModules(directory string) ([]string, bool, bool) {
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatal("Couldn't read directory")
	}

	var modules []string
	for _, file := range files {
		filePath := path.Join(directory, file.Name())

		isJsFile := strings.HasSuffix(filePath, ".js")
		if !isJsFile {
			continue
		}

		fileContentBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			continue
		}

		tree, err := js.Parse(parse.NewInputString(string(fileContentBytes)))
		if err != nil {
			return []string{}, false, false
		}

		for _, node := range tree.List {
			switch node := node.(type) {
			case *js.VarDecl:
				for _, binding := range node.List {
					parseExpression(binding.Default, &modules)
				}
			case *js.ExprStmt:
				parseExpression(node.Value, &modules)
			}
		}
	}

	var newRemoteModules []string
	canRemoveEnv := true
	for _, module := range modules {
		if checkIsLocalFile(module, directory) {
			continue
		}

		if checkIsBuiltin(module) {
			continue
		}

		if checkIsAlreadyInstalled(module, directory) {
			canRemoveEnv = false
			continue
		}

		newRemoteModules = append(newRemoteModules, module)
	}

	return newRemoteModules, canRemoveEnv, true
}
