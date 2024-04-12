package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/leonlatsch/fred/filewatcher"
	"github.com/leonlatsch/fred/sockets"
	"github.com/leonlatsch/fred/webserver"
)

var watcher *fsnotify.Watcher

func main() {
	go sockets.SetupSockets()
	go webserver.SetupWebServer()

	filewatcher.WatchFiles(func(name string) {
		log.Println("File changed. Refreshing...")
		if err := sockets.SendMessage(name); err != nil {
			log.Println(err)
		}
	})
}
