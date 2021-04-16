# Overview

Allows you to return a single deferred method that can be used to clean up more than one item in a block. This is really useful if you create lots of items that could throw errors when being created. When an error occurs, you want to clean up the items created and return the error, but if it works, you want to be able to pass to the caller a method they can use to clean up the items you created with defer in their scope.

You should be able to chain this up as well.

# Installing

`go get github.com/wojnosystems/go-defer-many`

# Using

```go
package main

import (
	"fmt"
	"github.com/wojnosystems/go-defer-many/deferMany"
	"ioutil"
	"os"
)

func main() {
	folders, cleanup, err := makeFolders()
	if err != nil {
		fmt.Panic("failed to make a folder!")
    }
	defer cleanup()

	// do stuff with folders
	_ = folders
}

func makeFolders() (folders []string, cleanup func(), err error) {
	// Create a new deferMany object that will allow us to aggregate tasks we want to defer until
	// this method is completed if there was an error, or skip the defer if there was no error
	deferrables := deferMany.New()
	defer deferrables.Defer()
	for i := 0; i < 200; i++ {
		var dir string
		dir, err = ioutil.TempDir("","")
		if err != nil {
			return
        }
        // TempDir created successfully, we want to be sure we clean this up!
        deferrables.Add(func(){
        	_ = os.RemoveAll(dir)
        })
	}
	
	// We're at the end of the method, we call "Return()" on deferrables. This causes deferrables.Defer to do nothing.
	// This prevents defer from removing your files as every tempDir was created successfully as
	// we want the caller to do something with those files and clean them up when they're done.
	// The function returned is all of the tasks you Add'ed.
	return folders, deferrables.Return(), nil
}
```
