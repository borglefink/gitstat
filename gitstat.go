package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/MichaelTJones/walk"
	"github.com/briandowns/spinner"
)

var (
	mutex    = sync.Mutex{}
	spinSet  = []string{"| ", "/ ", "- ", "\\ "}
	spin     = spinner.New(spinSet, 100*time.Millisecond)
	repodirs = make(sort.StringSlice, 0)
	startdir = "."
)

// init, makes sure we have a start directory
func init() {
	flag.Parse()
	var dir = flag.Arg(0)
	if dir != "" {
		startdir = dir
	}
}

// foreachEntry is called for each entry in the given directory
func foreachEntry(entryName string, f os.FileInfo, err error) error {
	if f == nil || !f.IsDir() || f.Name() != ".git" {
		return nil
	}

	// Getting the repo dir
	var dir = strings.Replace(entryName, string(os.PathSeparator)+".git", string(os.PathSeparator), 1)

	// edge case for current directory
	if dir == ".git" {
		dir = ""
	}

	mutex.Lock()
	repodirs = append(repodirs, dir)
	mutex.Unlock()

	return nil
}

// ensureDir makes sure that the given directory is not empty
// if empty it is the current directory
func ensureDir(dir string) string {
	if dir == "" {
		return "."
	}

	return dir
}

// printStatus runs "git status -s" for the given directory
func printStatus(workdir string) {
	var gitdir = workdir + ".git"

	var cmd = exec.Command("git", "--git-dir="+gitdir, "--work-tree="+ensureDir(workdir), "status", "-s")

	var r, w, _ = os.Pipe()
	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	cmd.Run()

	w.Close()
	var status, _ = ioutil.ReadAll(r)

	if len(bytes.TrimSpace(status)) > 0 {
		var workpath, err = filepath.Abs(ensureDir(workdir))
		if err != nil {
			workpath = workdir
		}

		fmt.Printf("\n--- %s ---\n", workpath)
		fmt.Printf("%v", string(status))
	}
}

// main
func main() {
	spin.Start()
	walk.Walk(startdir, foreachEntry)
	spin.Restart()
	spin.Stop()

	sort.Sort(repodirs)
	for _, wdir := range repodirs {
		printStatus(wdir)
	}
}
