package watch_files

import (
	"path/filepath"
	"strings"
	"time"

	. "github.com/cdvelop/output"
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

			if last_time, ok := last_actions[event.Name]; !ok || time.Since(last_time) > 3*time.Second {
				// Registrar la última acción y procesar el evento.
				last_actions[event.Name] = time.Now()

				if isDir(event.Name) {
					// fmt.Println("Folder Event:", event.Name)
				} else {
					// fmt.Println("File Event:", event.Name)

					extension := filepath.Ext(event.Name)

					switch extension {
					case ".css":
						PrintWarning("Compilando CSS..." + event.Name + "\n")
						err := action.BuildCSS()
						if err != nil {
							PrintError(err.Error())
						} else {
							err := action.Reload()
							if err != nil {
								PrintError(err.Error())
							}
						}
					case ".js":
						PrintWarning("Compilando JS..." + event.Name + "\n")
						err := action.BuildJS()
						if err != nil {
							PrintError(err.Error())
						} else {
							err := action.Reload()
							if err != nil {
								PrintError(err.Error())
							}
						}

					case ".html":
						PrintWarning("Compilando HTML..." + event.Name + "\n")
						err := action.BuildHTML()
						if err != nil {
							PrintError(err.Error())
						} else {
							err := action.Reload()
							if err != nil {
								PrintError(err.Error())
							}
						}

					case ".go":

						if strings.Contains(event.Name, "wasm") {
							PrintWarning("Compilando WASM..." + event.Name + "\n")
							err := action.BuildWASM()
							if err != nil {
								PrintError(err.Error())
							} else {

								err := action.Reload()
								if err != nil {
									PrintError(err.Error())
								}

							}
						} else {
							PrintWarning("Reiniciando APP..." + event.Name + "\n")
							err := action.Restart()
							if err != nil {
								PrintError(err.Error())
							} else {

								err := action.Reload()
								if err != nil {
									PrintError(err.Error())
								}
							}
						}

					}
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			PrintError(err.Error())
		}
	}

}
