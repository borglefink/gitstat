package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

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
