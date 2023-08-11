module github.com/cdvelop/watch_files

go 1.20

require github.com/fsnotify/fsnotify v1.6.0

require (
	github.com/cdvelop/output v0.0.2
	golang.org/x/sys v0.11.0 // indirect
)

replace github.com/cdvelop/output => ../output
