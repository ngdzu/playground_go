package security

import (
	"os"
	"path/filepath"
	"strings"
)

func Which() {
	file := "ls"
	path := os.Getenv("PATH")
	var foundit bool = false
	for _, directory := range strings.Split(path, ":") {
		fullpath := filepath.Join(directory, file)
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
