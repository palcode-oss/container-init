package main

import (
	"github.com/palcode-oss/container-init/languages"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("Not enough arguments")
	}

	language := os.Args[2]
	languages.Run(language)
}
