package main

import (
	"os"
	"path/filepath"
	"sort"
	"sync"
)

var (
	mutex    = sync.Mutex{}
	gitdirs  = make(sort.StringSlice, 0)
	startdir = "."
)

func init() {
	if len(os.Args[1:]) > 0 {
		startdir = filepath.Base(os.Args[1])
	}
}

func main() {
	findAllGitRepositories(startdir)

	for _, wdir := range gitdirs {
		printStatus(wdir)
	}
}
