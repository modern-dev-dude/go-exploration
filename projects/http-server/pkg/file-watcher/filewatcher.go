package filewatcher

// server/pkg/file-watcher

import (
	"log"
	"github.com/fsnotify/fsnotify"
)

func WatchFiles() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil{
		log.Fatal("Error creating watcher:", err)
	}
	// cleanup watcher every time main is restarted
	defer watcher.Close()

	err = watcher.Add(".")
	if err != nil {
		log.Fatal("Error adding watch:", err)
	}

	for {
		select {
		case event, ok := <- watcher.Events:
			if !ok {
				return
			}
			log.Println("File event: ", event)
		case err, ok :=  <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}

	}
}
