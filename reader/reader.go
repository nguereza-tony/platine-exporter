package reader

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Tail(path string, jobs chan<- []byte, offset *int64) {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	watcher.Add(path)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// Detect file update or create
				if event.Op&(fsnotify.Write|fsnotify.Create) != 0 {
					readFile(path, jobs, offset)
				}

				// Detect file remove or rename
				if event.Op&(fsnotify.Remove|fsnotify.Rename) != 0 {
					log.Println("File removed or renamed detected")
					time.Sleep(time.Second)

					fi, err := os.Stat(path)
					if err == nil {
						if fi.Size() < *offset {
							*offset = 0
						}
					}
				}

			case <-watcher.Errors:
			}
		}
	}()

	for {
		readFile(path, jobs, offset)
		time.Sleep(time.Second)
	}
}

func readFile(path string, jobs chan<- []byte, offset *int64) {
	log.Printf("Reading file: %s from offset %d", path, *offset)
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error reading file %s: %v", path, err)
		return
	}
	defer file.Close()

	fi, _ := file.Stat()
	if fi.Size() < *offset {
		log.Println("File rotation detected, resetting offset")
		*offset = 0
	}

	file.Seek(*offset, 0)

	buf := make([]byte, 8192)

	for {
		n, err := file.Read(buf)
		if n > 0 {
			*offset += int64(n)
			jobs <- append([]byte(nil), buf[:n]...)
		}

		if err != nil {
			break
		}
	}
}
