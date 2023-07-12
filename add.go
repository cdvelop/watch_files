package watch_files

func Add(c compiler, directory_folder_to_watch ...string) *ui {

	u := ui{
		directory_folders: directory_folder_to_watch,
		compiler:          c,
	}

	return &u
}
