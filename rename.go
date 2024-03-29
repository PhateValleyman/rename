package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	B = "\033[01;34m" // Blue color
	G = "\033[32m" // Green color
	R = "\033[31m" // Red color
	Y = "\033[01;33m" // Yellow color
	N = "\033[0m" // Reset colors
)

// showHelp displays usage information for the rename command.
func showHelp() {
	fmt.Printf("Usage: rename <%sfilename%s>\n", Y, N)
}

// showVersion displays version information for the rename command.
func showVersion() {
	fmt.Printf("%srename%s v1.1\n", Y, N)
	fmt.Printf("by %sPhateValleyman%s\n",Y, N)
	fmt.Printf("%sJonas.Ned@outlook.com%s\n", G, N)
}

// renameFile renames a file from oldFilename to newFilename.
// If newFilename already exists, it prompts the user for confirmation before overwriting.
func renameFile(oldFilename string, newFilename string, defaultFilename string) error {
	// If newFilename is empty, set it to defaultFilename
	if newFilename == "" {
		newFilename = defaultFilename
	}

	// Check if newFilename already exists
	_, err := os.Stat(newFilename)
	if err == nil {
		fmt.Printf("\"%s%s%s\"", Y, newFilename, N)
		fmt.Print(" already exists. Do you want to overwrite it? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		overwrite, _ := reader.ReadString('\n')
		overwrite = strings.TrimSpace(overwrite)
		switch strings.ToLower(overwrite) {
		case "y", "yes":
			// If user confirms, rename the file
			if err := os.Rename(oldFilename, newFilename); err != nil {
				fmt.Printf("Failed to rename file \"%s%s%s\"\n", R, oldFilename, N)
				return err
			}
			fmt.Printf("File \"%s%s%s\" renamed to \"%s%s%s\"\n", Y, oldFilename, N, G, newFilename, N)
		default:
			// If user cancels, abort renaming
			fmt.Println("Rename aborted...")
		}
	} else {
		// If newFilename doesn't exist, rename the file
		if err := os.Rename(oldFilename, newFilename); err != nil {
			fmt.Printf("Failed to rename file \"%s%s%s\"\n", R, oldFilename, N)
			return err
		}
		fmt.Printf("File \"%s%s%s\" renamed to \"%s%s%s\"\n", Y, oldFilename, N, G, newFilename, N)
	}
	return nil
}

// processArguments processes the command line arguments passed to the rename command.
// It validates the arguments, prompts the user for input, and calls renameFile to perform the renaming operation.
func processArguments(args []string) error {
	// Check if there are enough arguments
	if len(args) < 1 {
		showHelp()
		return fmt.Errorf("missing arguments")
	}
	// Get the old filename from the arguments
	oldFilename := args[0]
	// Check if the old filename exists
	if _, err := os.Stat(oldFilename); err != nil {
		showHelp()
		return fmt.Errorf("invalid filename")
	}
	// Prompt the user to enter a new filename
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%sEnter new filename (default: %s): %s", Y, oldFilename, N)
	newFilename, _ := reader.ReadString('\n')
	newFilename = strings.TrimSpace(newFilename)
	// If the new filename is empty, set it to the old filename
	if newFilename == "" {
		newFilename = oldFilename
	}
	// Call renameFile to perform the renaming operation
	return renameFile(oldFilename, newFilename, oldFilename)
}

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 2 {
		showHelp()
		return
	}
	// Process the command specified by the first argument
	switch os.Args[1] {
	case "--help", "-h", "-?":
		showHelp()
	case "--version", "-v":
		showVersion()
	default:
		if err := processArguments(os.Args[1:]); err != nil {
			fmt.Println(err)
		}
	}
}
