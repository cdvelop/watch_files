package watch_files

func Add(c compiler, d reload, t restart, directories_to_watch map[string]struct{}) *WatchFiles {

	w := WatchFiles{
		watch_dir_folders: directories_to_watch,
		compiler:          c,
		reload:            d,
		restart:           t,
	}

	return &w
}
