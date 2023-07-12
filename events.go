package watch_files

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func (u ui) watchEvents(watcher *fsnotify.Watcher) {
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
						u.BuildCSS()
						// RELOADED HERE
						showMessage("reload_browser")
					case ".js":
						fmt.Println("Compilando JS...", event.Name)
						u.BuildJS()
						// RELOADED HERE
						showMessage("reload_browser")
					case ".html":
						fmt.Println("Compilando HTML...", event.Name)
						u.BuildHTML()
						// RELOADED HERE
						showMessage("reload_browser")
					case ".go":

						if strings.Contains(event.Name, "wasm") {
							fmt.Println("Compilando WASM...", event.Name)
							u.BuildWASM()
							// RELOADED HERE

							showMessage("reload_browser")
						} else {
							showMessage("restart_app")
							time.Sleep(10 * time.Millisecond)
							os.Exit(0)

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
