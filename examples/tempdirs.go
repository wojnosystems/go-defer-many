package main

import (
	"fmt"
	"github.com/wojnosystems/go-defer-many/deferMany"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dirs, cleanup, err := makeTmpDirs()
	if err != nil {
		log.Panic("failed to make a directory!")
	}
	defer cleanup()

	// do stuff with dirs
	for _, dir := range dirs {
		fmt.Println("pretend to work with: ", dir)
	}

	// when main exits, all directories created will be deleted due to defer cleanup()
}

func makeTmpDirs() (dirs []string, cleanup func(), err error) {
	// Create a new deferMany object that will allow us to aggregate tasks we want to defer until
	// this method is completed if there was an error, or skip the defer if there was no error
	deferred := deferMany.New()
	defer deferred.Defer()
	for i := 0; i < 5; i++ {
		var dir string
		dir, err = ioutil.TempDir("", "")
		if err != nil {
			// if some error occurs, any directories created so far will be removed before this method returns
			return
		}
		dirs = append(dirs, dir)
		// TempDir created successfully, we want to be sure we clean this up!
		deferred.Add(func() {
			fmt.Println("deleting! ", dir)
			_ = os.RemoveAll(dir)
		})
	}

	// We're at the end of the method, we call "Return()" on deferred. This causes deferred.Defer to do nothing.
	// This prevents defer from removing your files as every tempDir was created successfully as
	// we want the caller to do something with those files and clean them up when they're done.
	// The function returned is all of the tasks you Add'ed.
	return dirs, deferred.Return(), nil
}
