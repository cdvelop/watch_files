package watch_files

type WatchFiles struct {
	DIRECTORY_FOLDERS []string // ej: "modules", "ui\\theme"
	compiler
}

type compiler interface {
	BuildHTML()
	BuildCSS()
	BuildJS()
	BuildWASM()
}
