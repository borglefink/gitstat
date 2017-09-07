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
	mutex      = sync.Mutex{}
	spinSet    = []string{"| ", "/ ", "- ", "\\ "}
	spin       = spinner.New(spinSet, 100*time.Millisecond)
	repodirs   = make(sort.StringSlice, 0)
	startdir   = "."
	alwaysshow = flag.Bool(
		"a",
		false,
		"show directory even if there is no status",
	)
	long = flag.Bool(
		"l",
		false,
		"show normal git status",
	)
	ignored = flag.Bool(
		"i",
		false,
		"show 'git status --ignored')",
	)
	help = flag.Bool(
		"?",
		false,
		"shows this help",
	)
)

// init, makes sure we have a start directory
func init() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		usage()
		os.Exit(0)
	}

	var dir = flag.Arg(0)
	if dir != "" {
		startdir = dir
	}
}

// usage
func usage() {
	fmt.Printf("\nGITSTAT (C) Copyright 2017 Erlend Johannessen\n")
	fmt.Printf("Finds all git repositories below the given path, and for each repository runs \"git status -s\".\n")
	fmt.Printf("Usage: gitstat [options] [dirname] \n")
	flag.PrintDefaults()
	fmt.Printf("\n")
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

	var dirpath, patherr = filepath.Abs(ensureDir(dir))
	if patherr != nil {
		return nil
	}

	mutex.Lock()
	repodirs = append(repodirs, dirpath)
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
	var gitdir = workdir + string(os.PathSeparator) + ".git"

	var args = []string{
		"--git-dir=" + gitdir,
		"--work-tree=" + ensureDir(workdir),
		"status",
	}

	if !*long {
		args = append(args, "-s")
	}

	if *ignored {
		args = append(args, "--ignored")
	}

	var cmd = exec.Command("git", args...)

	var r, w, _ = os.Pipe()
	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	cmd.Run()

	w.Close()
	var status, _ = ioutil.ReadAll(r)

	if *alwaysshow || len(bytes.TrimSpace(status)) > 0 {
		fmt.Printf("----- %s -----\n", workdir)
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
