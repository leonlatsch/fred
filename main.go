package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/leonlatsch/fred/config"
	"github.com/leonlatsch/fred/filewatcher"
	"github.com/leonlatsch/fred/sockets"
	"github.com/leonlatsch/fred/webserver"
)

var watcher *fsnotify.Watcher

func main() {
	conf := config.GetConfig()

	go sockets.SetupSockets()
	go webserver.SetupWebServer(conf.Port)

	filewatcher.WatchFiles(func(name string) {
		log.Println("File changed. Refreshing...")
		if err := sockets.SendMessage(name); err != nil {
			log.Println(err)
		}
	})
}
