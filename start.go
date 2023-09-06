package watch_files

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// La función principal es donde se crea el observador para monitorear los cambios en los archivos y directorios.
// En esta función, también configuraremos los filtros para los tipos de archivo que queremos observar.
func (w WatchFiles) DevFileWatcherSTART(wg *sync.WaitGroup) {
	defer wg.Done()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer watcher.Close()

	go w.watchEvents(watcher)
	reg := map[string]struct{}{}
	for folder := range w.watch_dir_folders {
		// fmt.Println("carpeta registrada:", folder)
		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && !w.Contain(path) {
				if _, exist := reg[path]; !exist {

					watcher.Add(path)

					reg[path] = struct{}{}

					// fmt.Println("** NEW PATH:", path)
				}
			}
			return nil
		})
	}

	fmt.Println("Escuchando Eventos UI ...")
	select {}
}
