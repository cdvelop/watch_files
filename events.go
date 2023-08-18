package watch_files

import (
	"path/filepath"
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
				var err error
				// Registrar la última acción y procesar el evento.
				last_actions[event.Name] = time.Now()

				if isDir(event.Name) {
					// fmt.Println("Folder Event:", event.Name)
				} else {
					// fmt.Println("File Event:", event.Name)

					extension := filepath.Ext(event.Name)

					switch extension {
					case ".css":
						err = action.BuildCSS(event.Name)
						if err != nil {
							PrintError(err.Error())
						}
						reload(err)

					case ".js":
						err = action.BuildJS(event.Name)
						if err != nil {
							PrintError(err.Error())
						}
						reload(err)

					case ".html":
						err = action.BuildHTML(event.Name)
						if err != nil {
							PrintError(err.Error())
						}
						reload(err)

					case ".go":

						err = action.BuildWASM(event.Name)
						if err != nil {
							PrintError(err.Error())
						} else {

							err = action.Restart(event.Name)
							if err != nil {
								PrintError(err.Error())
							}
						}

						reload(err)
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

func reload(err error) {
	if err == nil {
		err := action.Reload()
		if err != nil {
			PrintError(err.Error())
		}
	}
}
