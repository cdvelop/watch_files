package watch_files

type WatchFiles struct {
	watch_dir_folders []string // ej: "modules", "ui\\theme"
	compiler
	reload
	restart
}

type compiler interface {
	BuildHTML() error
	BuildCSS() error
	BuildJS() error
	BuildWASM() error
}

type reload interface {
	Reload()
}

type restart interface {
	Restart() error
}
