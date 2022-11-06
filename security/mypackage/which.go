package mypackage

import (
	"os"
	"strings"
)

func which() {
	file := "which"
	path := os.Getenv("PATH")
	var foundit bool = false
	for _, directory := range strings.Split(path, ":") {
		fullpath := directory + file
		fileinfo, err := os.Lstat(fullpath)
		if err == nil {
			mode := fileinfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					foundit = true
				}
			}
		}
	}

	if !foundit {
		os.Exit(1)
	}
}
