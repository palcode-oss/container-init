package main

import (
	"github.com/palcode-oss/runner-common/languages"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough arguments")
	}

	language := os.Args[2]
	languages.Run(language)
}
