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
	BuildHTML(event_name string) error
	BuildCSS(event_name string) error
	BuildJS(event_name string) error
	BuildWASM(event_name string) error
}

type reloadAPP interface {
	Reload() error
}

type restartAPP interface {
	Restart(event_name string) error
}
