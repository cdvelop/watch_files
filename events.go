package watch_files

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func (w WatchFiles) watchEvents(watcher *fsnotify.Watcher) {
	// defer wg.Done()
	last_actions := make(map[string]time.Time)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if last_time, ok := last_actions[event.Name]; !ok || time.Since(last_time) > 2*time.Second {
				// Registrar la última acción y procesar el evento.
				last_actions[event.Name] = time.Now()

				if isDir(event.Name) {
					// fmt.Println("Folder Event:", event.Name)
				} else {
					// fmt.Println("File Event:", event.Name)

					extension := filepath.Ext(event.Name)

					switch extension {
					case ".css":
						fmt.Println("Compilando CSS...", event.Name)
						err := w.BuildCSS()
						if err != nil {
							log.Println(err)
						} else {
							w.Reload()
						}
					case ".js":
						fmt.Println("Compilando JS...", event.Name)
						err := w.BuildJS()
						if err != nil {
							log.Println(err)
						} else {
							w.Reload()
						}

					case ".html":
						fmt.Println("Compilando HTML...", event.Name)
						err := w.BuildHTML()
						if err != nil {
							log.Println(err)
						} else {
							w.Reload()
						}

					case ".go":

						if strings.Contains(event.Name, "wasm") {
							fmt.Println("Compilando WASM...", event.Name)
							err := w.BuildWASM()
							if err != nil {
								log.Println(err)
							} else {
								w.Reload()
							}
						} else {

							err := w.Restart()
							if err != nil {
								log.Println("Restart error: ", err)
							}

							// showMessage("restart_app")
							// time.Sleep(10 * time.Millisecond)
							// os.Exit(0)

						}

					}
				}

			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}

}
