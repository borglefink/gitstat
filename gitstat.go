package main

import (
	"bytes"
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
	spinSet  = []string{"| ", "/ ", "- ", "\\ "}
	spin     = spinner.New(spinSet, 100*time.Millisecond)
	repodirs = make(sort.StringSlice, 0)
	startdir = "."
	mutex    = sync.Mutex{}
)

func init() {
	if len(os.Args[1:]) > 0 {
		startdir = filepath.Base(os.Args[1])
	}
}

func foreachEntry(entryName string, f os.FileInfo, err error) error {
	if f == nil {
		return nil
	}

	if !f.IsDir() {
		return nil
	}

	if f.Name() != ".git" {
		return nil
	}

	// Getting the working dir
	var dir = strings.Replace(entryName, ".git", "", 1)

	mutex.Lock()
	repodirs = append(repodirs, dir)
	mutex.Unlock()

	return nil
}

func printStatus(workdir string) {
	var gitdir = workdir + ".git"
	var cmd = exec.Command("git", "--git-dir="+gitdir, "--work-tree="+workdir, "status", "-s")

	var r, w, _ = os.Pipe()
	cmd.Stdout = w
	cmd.Stderr = os.Stderr
	cmd.Run()
	w.Close()

	var status, _ = ioutil.ReadAll(r)

	if len(bytes.TrimSpace(status)) > 0 {
		var workpath, err = filepath.Abs(workdir)
		if err != nil {
			workpath = workdir
		}

		fmt.Println(workpath)
		fmt.Println(string(status))
	}
}

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
