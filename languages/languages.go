package languages

import (
	"github.com/palcode-oss/runner-common/languages/node"
	"github.com/palcode-oss/runner-common/languages/python"
)

var languages = map[string]func(){
	"python": python.Run,
	"node":   node.Run,
}

func Run(language string) {
	runner, ok := languages[language]
	if !ok {
		panic("Language not found")
	} else {
		runner()
	}
}
