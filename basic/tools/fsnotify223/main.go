package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

var (
	watcher *fsnotify.Watcher
)

func main() {
	fmt.Printf("%#v\n", watcher)
	if watcher == nil {
		// error!!!!!!!!!!!!!!!!!
		// if one var is undefined
		// then all of them are new vars
		watcher, err := fsnotify.NewWatcher()
		fmt.Printf("%#v\n", watcher)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	// watcher == nil here
	fmt.Printf("%#v\n", watcher)
	watcher.Add("test")
}

func f() {
	fmt.Printf("%#v\n", watcher)
	if watcher == nil {
		// error!!!!!!!!!!!!!!!!!
		var err error
		watcher, err = fsnotify.NewWatcher()
		fmt.Printf("%#v\n", watcher)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Printf("%#v\n", watcher)
	watcher.Add("test")
}
