package watch_files

func Add(c compiler, d reload, t restart, directory_folder_to_watch ...string) *WatchFiles {

	w := WatchFiles{
		watch_dir_folders: directory_folder_to_watch,
		compiler:          c,
		reload:            d,
		restart:           t,
	}

	return &w
}
