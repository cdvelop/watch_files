package watch_files

import "github.com/cdvelop/gotools"

var action app

func Add(com compilerAPP, rel reloadAPP, res restartAPP, directories_to_watch map[string]struct{}) *WatchFiles {

	if com == nil {
		gotools.ShowErrorAndExit("error compilador nulo")
	}
	if rel == nil {
		gotools.ShowErrorAndExit("error Manejador app reload nulo")
	}

	if res == nil {
		gotools.ShowErrorAndExit("error Manejador app restart nulo")
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
