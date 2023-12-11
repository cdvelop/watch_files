package watch_files

import (
	"path/filepath"
	"time"

	. "github.com/cdvelop/output"
	"github.com/fsnotify/fsnotify"
)

func (w WatchFiles) watchEvents(watcher *fsnotify.Watcher) {
	last_actions := make(map[string]time.Time)
	reloadTimer := time.NewTimer(0)
	reloadTimer.Stop()

	restarTimer := time.NewTimer(0)
	restarTimer.Stop()

	var wait = 50 * time.Millisecond

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if last_time, ok := last_actions[event.Name]; !ok || time.Since(last_time) > 1*time.Second {
				var err string

				// Restablece el temporizador de recarga
				reloadTimer.Stop()

				if isDir(event.Name) {
					// fmt.Println("Folder Event:", event.Name)
				} else {
					// fmt.Println("File Event:", event.Name)

					extension := filepath.Ext(event.Name)

					switch extension {
					case ".css":
						err = action.BuildCSS(event.Name)
						if err == "" {
							reloadTimer.Reset(wait)
						}
					case ".js":
						err = action.BuildJS(event.Name)
						if err == "" {
							reloadTimer.Reset(wait)
						}
					case ".html":
						err = action.BuildHTML(event.Name)
						if err == "" {
							reloadTimer.Reset(wait)
						}

					case ".go":
						restarTimer.Stop()
						restarTimer.Reset(wait)

					}

					if err != "" {
						PrintError(err)
					}
				}

				// Registrar la última acción
				last_actions[event.Name] = time.Now()

			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			PrintError(err.Error())

		case <-restarTimer.C:

			err := action.BuildWASM(time.Now().Format("15:04:05"))
			if err == "" {
				err = action.Restart(time.Now().Format("15:04:05"))
				if err == "" {
					reload()
				}
			}

			if err != "" {
				PrintError(err)
			}

		case <-reloadTimer.C:
			// El temporizador de recarga ha expirado, ejecuta reload()
			reload()

		}
	}
}

func reload() {
	err := action.Reload()
	if err != "" {
		PrintError(err)
	}
}
