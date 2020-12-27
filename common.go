package common

import (
	"fmt"
	"os"
	"path"
)

func ModuleMessage() {
	lockFilePath := path.Join(
		"/usr/src/app", ".module_info_lock",
	)

	_, err := os.Stat(lockFilePath)
	if os.IsNotExist(err) {
		file, _ := os.Create(lockFilePath)
		_ = file.Close()

		fmt.Println("It looks like you're using 3rd-party modules in your code.")
		fmt.Println("I'll install them for you now; please be patient.")
		fmt.Println("If you think I've detected this incorrectly, please email contact@palcode.dev.")
	}
}
