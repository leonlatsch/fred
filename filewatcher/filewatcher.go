package filewatcher

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func WatchFiles(onChange func(name string)) {
	watcher, _ = fsnotify.NewWatcher()
	watchDirs(".")

	defer watcher.Close()

	for {
		event := <-watcher.Events

		switch event.Op {
		case fsnotify.Create:
			stat, err := os.Stat(event.Name)
			if err == nil && stat.IsDir() {
				log.Println("Created " + event.Name)
				watcher.Add(event.Name)
			}
		case fsnotify.Remove:
			log.Println("Deleted " + event.Name)
			watcher.Remove(event.Name)
		}

		onChange(event.Name)
	}
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
