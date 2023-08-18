package watch_files

import . "github.com/cdvelop/output"

var action app

func Add(com compilerAPP, rel reloadAPP, res restartAPP, directories_to_watch map[string]struct{}, rebuild_when_dir_contains ...string) *WatchFiles {

	if com == nil {
		ShowErrorAndExit("error compilador nulo")
	}
	if rel == nil {
		ShowErrorAndExit("error Manejador app reload nulo")
	}

	if res == nil {
		ShowErrorAndExit("error Manejador app restart nulo")
	}

	action = app{
		compilerAPP: com,
		reloadAPP:   rel,
		restartAPP:  res,
	}

	w := WatchFiles{
		watch_dir_folders: directories_to_watch,
	}

	return &w
}
