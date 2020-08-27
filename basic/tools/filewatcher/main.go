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
					fmt.Printf("%#v\n", event)
					if cb, ok := gCb[event.Name]; ok && event.Op&fsnotify.Remove == fsnotify.Remove {
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
	fmt.Println(path)
	absPath, err := filepath.Abs(path)
	fmt.Println(absPath)
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
		gCloseCh = nil
		gWatcher.Close()
		gWatcher = nil
		gCb = nil
	}
}

func main() {
	called := false
	f, _ := os.Create("testfile")
	f.Close()

	err := WatchFile("testfile", func() {
		fmt.Println("Trigger!")
		called = true
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := os.Remove("testfile"); err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(time.Second)

	if called == true {
		fmt.Println("Success")
	} else {
		fmt.Println("Not ok")
	}
}
