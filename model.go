package watch_files

type ui struct {
	directory_folders []string // ej: "modules", "ui\\theme"
	compiler
}

type compiler interface {
	BuildHTML()
	BuildCSS()
	BuildJS()
	BuildWASM()
}
