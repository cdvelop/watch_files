package watch_files

func Add(c compiler, directory_folder_to_watch ...string) *WatchFiles {

	w := WatchFiles{
		DIRECTORY_FOLDERS: directory_folder_to_watch,
		compiler:          c,
	}

	return &w
}
