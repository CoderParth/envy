package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// Envs is a struct that holds a slice of string called "vars"
// and a mutex to synchronize the access to the slice.
//
// During the recursive walk of the file directory, all the
// extracted env variables are stored in its instance
// which has been declared few lines below as "e".
type Envs struct {
	vars []string
	mu   sync.Mutex
}

var (
	e  Envs
	wg sync.WaitGroup // wg waits for all the goroutines to finish their executions.
)

// filesToExclude is a map that contains the list of filenames, and
// file dirs that needs to be skipped, while recursively walking
// the file tree rooted at root.
var filesToExclude = map[string]int{
	".git":                   0,
	".gitignore":             0,
	"package-lock.json":      0,
	"logs":                   0,
	"node_modules":           0,
	"dist":                   0,
	"build":                  0,
	"out":                    0,
	".env":                   0,
	".env.local":             0,
	".env.development.local": 0,
	".env.test.local":        0,
	".env.production.local":  0,
	".env.example":           0,
	".vscode":                0,
	".idea":                  0,
	"main.go":                0,
	"README.md":              0,
	"envy":                   0,
}

// scanPath is used as a WalkFunc for the filepath.Walk() function
// to visit each file or directory.
//
// It is called for each file and dir in the file root.
//
// For each file or dir, the func checks if the file name/dir name is present in the
// globally declared map called "filesToExclude". If the file/dir exists in the
// map, the file or dir is skipped.
//
// Only if the file/dir is not present in the map, readFileAndExtractEnvs() is called
// which takes in the filename of string type as an argument.
func scanPath(path string, info os.FileInfo, err error) error {
	// Get the name of the current file path
	_, name := filepath.Split(path)

	// Check if the path exists in "filesToExclude"
	if _, ok := filesToExclude[name]; ok {
		// If the current path is a directory, which is present in filesToExclude,
		// skip the curr dir.
		if info.IsDir() {
			fmt.Printf("Skipping dir: %s\n", path)
			return filepath.SkipDir
		}
		// Since the current path is a file and is present in filesToExclude,
		// skip the curr file
		fmt.Printf("Skipping file : %s\n", path)
		return nil
	}

	// Else, if it not a dir, increment the counter in the WaitGroup, and call readFileAndExtractEnvs()
	if !info.IsDir() {
		wg.Add(1)
		go readFileAndExtractEnvs(path)
	}

	return nil
}

// readFileAndExtractEnvs takes in the file name as an argument.
// It opens the file provied in the argument, scans the file,
// iterates through each line, and looks for lines that contain "process.env".
//
// And from such lines, env variable names are extracted and
// appended to the globally declared instance of "Envs" called "e".
//
// In case of errors, a fatal log is printed, and the func moves on to the next file/dir.
func readFileAndExtractEnvs(file string) {
	defer wg.Done()
	readfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer readfile.Close()

	fmt.Printf("Reading file: %s\n", file)
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	// Iterate through each line and find string pattern "process.env".
	// From the found patterns, extract variable names and
	// store in a globally declared slice called "envs"
	for fileScanner.Scan() {
		currLine := fileScanner.Text()
		hasEnvVariable := strings.Contains(currLine, "process.env.")
		if hasEnvVariable {
			vars := strings.Split(currLine, "process.env.")
			fmt.Printf("Vars : %v\n", vars)
			for i, v := range vars {
				// First index contains empty string, so skip.
				if i == 0 {
					continue
				}
				// Use a regular expression to match only alphanumeric characters and underscores
				re := regexp.MustCompile(`[\w_]+`)
				varName := re.FindString(v)

				e.mu.Lock()
				e.vars = append(e.vars, varName)
				e.mu.Unlock()
			}
		}
	}
}

// createEnvExample creates a new "env.example" file,
// iterates over all the environment variables stored in
// globally declared instance of "envs" struct called "e",
// then writes each env to the ".env.example" file in
// "variableName=" format with new line per each variable.
//
// If ".env.example" already exists then the file is overwritten.
func createEnvExample() {
	file, err := os.Create(".env.example")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	for _, v := range e.vars {
		fmt.Println(v)
		v := v + "=" + "\n"
		if _, err := file.WriteString(v); err != nil {
			log.Fatal(err)
		}
	}

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Created .env.example file in the root")
}

func main() {
	if err := filepath.Walk(".", scanPath); err != nil {
		fmt.Printf("Error scanning the directory %v:\n", err)
	}

	wg.Wait()
	fmt.Println("Finished walking the path")

	createEnvExample()
}
