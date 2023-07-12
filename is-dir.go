package watch_files

import (
	"log"
	"os"
)

func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return false
	}
	return fi.Mode().IsDir()
}
