package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	dir, _ := getDir()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				fmt.Println("Event: ", event)

				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("Created file: ", event.Name)
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Modified file: ", event.Name)
				}

				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("Removed file: ", event.Name)
				}

				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Println("Renamed file: ", event.Name)
				}

				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					fmt.Println("Chmod file: ", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error: ", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		panic(err)
	}

	<-make(chan struct{})
}

func getDir() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter directory to watch: ")

	scanner.Scan()

	fmt.Println("Watch dir: ", scanner.Text())
	fmt.Println("----------------------------------")

	return scanner.Text(), scanner.Err()
}
