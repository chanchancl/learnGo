package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type FileChangeCallback func()

var (
	gWatcher *fsnotify.Watcher
	gMutex   sync.Mutex
	gCb      map[string]FileChangeCallback
	gCloseCh chan struct{}
)

// WatchFile only watch the REMOVE operation
func WatchFile(path string, cb FileChangeCallback) error {
	gMutex.Lock()
	defer gMutex.Unlock()
	if gWatcher == nil {
		// Init watcher
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return err
		}
		gWatcher = watcher
		gCloseCh = make(chan struct{})
		gCb = make(map[string]FileChangeCallback)
		go func() {
			for {
				if gWatcher == nil {
					return
				}
				select {
				case event, ok := <-gWatcher.Events:
					if !ok {
						return
					}
					fmt.Println("Received event")
					fmt.Printf("%#v\n", event)
					if cb, ok := gCb[event.Name]; ok {
						cb()
					}
				case _, ok := <-gWatcher.Errors:
					if !ok {
						return
					}
				case <-gCloseCh:
					return
				}
			}
		}()
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	err = gWatcher.Add(absPath)
	if err != nil {
		return err
	}
	gCb[absPath] = cb
	return nil
}

// Stop stops all of monitors
func Stop() {
	gMutex.Lock()
	defer gMutex.Unlock()
	if gWatcher != nil {
		gCloseCh <- struct{}{}
		close(gCloseCh)
		gWatcher.Close()
		gCloseCh = nil
		gWatcher = nil
		gCb = nil
	}
}

func main() {
	called := make(chan struct{})
	f, _ := os.Create("testfile")
	f.Close()
	defer os.Remove("teestfile")
	err := WatchFile("testfile", func() {
		called <- struct{}{}
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Remove("testfile")
	select {
	case <-time.Tick(time.Second):
		fmt.Println("Failed!")
	case <-called:
		fmt.Println("Success")
	}
}
