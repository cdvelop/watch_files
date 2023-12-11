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
	BuildHTML(event_name string) (err string)
	BuildCSS(event_name string) (err string)
	BuildJS(event_name string) (err string)
	BuildWASM(event_name string) (err string)
}

type reloadAPP interface {
	Reload() (err string)
}

type restartAPP interface {
	Restart(event_name string) (err string)
}
