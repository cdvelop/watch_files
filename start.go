package watch_files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// La función principal es donde se crea el observador para monitorear los cambios en los archivos y directorios.
// En esta función, también configuraremos los filtros para los tipos de archivo que queremos observar.
func (w WatchFiles) DevFileWatcherSTART() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer watcher.Close()

	go w.watchEvents(watcher)

	for _, folder := range w.DIRECTORY_FOLDERS {

		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && !contain(path) {
				// fmt.Println(path)
				watcher.Add(path)
			}
			return nil
		})
	}

	fmt.Println("Escuchando Eventos UI ...")
	select {}
}
