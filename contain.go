package watch_files

import "strings"

func contain(path string) bool {
	var no_add = []string{".git", ".vscode", "built"}

	for _, value := range no_add {
		if strings.Contains(path, value) {
			return true
		}
	}

	return false
}
