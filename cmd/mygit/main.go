// Package main implements a minimal Git command-line tool.
//
// It is the entry point for the "Build Your Own Git" CodeCrafters challenge.
// Currently supported commands:
//   - init: Initialise a new Git repository in the current directory.
package main

import (
	"fmt"
	"os"
)

// initRepo initializes a new Git repository in the current working directory.
// It creates the required directory structure under .git:
//   - .git/
//   - .git/objects/
//   - .git/refs/
//
// It also writes the default HEAD file pointing to refs/heads/main.
// Any error encountered while creating directories or writing the HEAD file is
// printed to stderr; the function does not abort on those errors so that partial
// setup is still reported to the caller.
func initRepo() {
	// Create the three directories that every Git repository requires.
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	// Write the HEAD file. A fresh repository points to the default branch (main).
	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}

	fmt.Println("Initialized git directory")
}

// main is the program entry point.
//
// It reads the sub-command from os.Args and dispatches to the appropriate
// handler function.  An unknown sub-command is reported to stderr and the
// process exits with status 1.
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		initRepo()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
