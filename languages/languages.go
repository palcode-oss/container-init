package languages

import (
	"github.com/palcode-oss/container-init/languages/bash"
	"github.com/palcode-oss/container-init/languages/cpp"
	_go "github.com/palcode-oss/container-init/languages/go"
	"github.com/palcode-oss/container-init/languages/java"
	"github.com/palcode-oss/container-init/languages/node"
	"github.com/palcode-oss/container-init/languages/php"
	"github.com/palcode-oss/container-init/languages/prolog"
	"github.com/palcode-oss/container-init/languages/python"
)

var languages = map[string]func(){
	"python": python.Run,
	"nodejs": node.Run,
	"bash":   bash.Run,
	"java":   java.Run,
	"prolog": prolog.Run,
	"go":     _go.Run,
	"cpp":    cpp.Run,
	"php":    php.Run,
}

func Run(language string) {
	runner, ok := languages[language]
	if !ok {
		panic("Language not found")
	} else {
		runner()
	}
}
