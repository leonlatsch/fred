package filewatcher

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func WatchFiles(onChange func(name string)) {
	createAndWatch()

	defer watcher.Close()

	for {
		event := <-watcher.Events
		createAndWatch()

		onChange(event.Name)
	}
}

func createAndWatch() {
	if watcher != nil {
		watcher.Close()
	}

	watcher, _ = fsnotify.NewWatcher()
	watchDirs(".")
}

func watchDirs(path string) {
	watcher.Add(path)
	filepath.WalkDir(path, func(childPath string, d fs.DirEntry, err error) error {
		if strings.HasPrefix(childPath, ".") {
			return nil
		}

		if d.IsDir() {
			watcher.Add(childPath)
		}
		return nil
	})
}
