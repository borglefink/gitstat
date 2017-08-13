package main

// func makeAndRunFile() {
// 	// create content
// 	var batchFile = make([]string, 0)
// 	batchFile = append(batchFile, "@echo off")
// 	for _, item := range gitdirs {
// 		batchFile = append(batchFile, "cd "+item)
// 		batchFile = append(batchFile, "echo ----- "+item)
// 		batchFile = append(batchFile, "git status -s")
// 	}

// 	// create
// 	var fileContents = strings.Join(batchFile, newline)
// 	var err = ioutil.WriteFile(BatchFilename, []byte(fileContents), 0644)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(-1)
// 	}

// 	// execute
// 	var cmd = exec.Command(BatchFilename)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	err = cmd.Run()

// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(-1)
// 	}

// 	// remove
// 	err = os.Remove(BatchFilename)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		os.Exit(-1)
// 	}
// }
