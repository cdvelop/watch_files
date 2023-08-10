package watch_files

type WatchFiles struct {
	watch_dir_folders map[string]struct{} // ej: "modules", "ui\\theme"
}

type app struct {
	compilerAPP
	reloadAPP
	restartAPP
}

type compilerAPP interface {
	BuildHTML() error
	BuildCSS() error
	BuildJS() error
	BuildWASM() error
}

type reloadAPP interface {
	Reload() error
}

type restartAPP interface {
	Restart() error
}
