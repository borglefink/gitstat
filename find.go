package main

import (
	"os"
	"sort"
	"strings"

	"github.com/MichaelTJones/walk"
)

func foreachEntry(entryName string, f os.FileInfo, err error) error {
	if f == nil {
		// Just ignore file errors from system
		return nil
	}

	if !f.IsDir() {
		// Ignore files
		return nil
	}

	if f.Name() != ".git" {
		// Ignore everything except .git repositories
		return nil
	}

	var dir = strings.Replace(entryName, ".git", "", 1)

	// if not found, just ignore
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}

	mutex.Lock()
	gitdirs = append(gitdirs, dir)
	mutex.Unlock()

	return nil
}

func findAllGitRepositories(dir string) {
	walk.Walk(dir, foreachEntry)
	sort.Sort(gitdirs)
}
