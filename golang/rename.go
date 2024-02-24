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

// Function to rename a file
func renameFile(oldFilename string, newFilename string, defaultFilename string) error {
	if newFilename == "" {
		newFilename = defaultFilename
	}
	_, err := os.Stat(newFilename)
	if err == nil {
		fmt.Printf("\"%s%s%s\"", Y, newFilename, N)
		fmt.Print(" already exists. Do you want to overwrite/rename/no? (o/r/n): ")
		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		switch strings.ToLower(option) {
		case "o", "overwrite":
			if err := os.Rename(oldFilename, newFilename); err != nil {
				fmt.Printf("Failed to rename file \"%s%s%s\"\n", R, oldFilename, N)
				return err
			}
			fmt.Printf("File \"%s%s%s\" renamed to \"%s%s%s\"\n", Y, oldFilename, N, G, newFilename, N)
		case "r", "rename":
			fmt.Printf("Enter new name for \"%s%s%s\": ", Y, newFilename, N)
			newFilename, _ = reader.ReadString('\n')
			newFilename = strings.TrimSpace(newFilename)
			return renameFile(oldFilename, newFilename, defaultFilename)
		default:
			fmt.Println("Rename aborted...")
			return nil
		}
	} else {
		if err := os.Rename(oldFilename, newFilename); err != nil {
			fmt.Printf("Failed to rename file \"%s%s%s\"\n", R, oldFilename, N)
			return err
		}
		fmt.Printf("File \"%s%s%s\" renamed to \"%s%s%s\"\n", Y, oldFilename, N, G, newFilename, N)
	}
	return nil
}

// Function to process command line arguments
func processArguments(args []string) error {
	if len(args) < 1 {
		showHelp()
		return fmt.Errorf("missing arguments")
	}
	oldFilename := args[0]
	if _, err := os.Stat(oldFilename); err != nil {
		showHelp()
		return fmt.Errorf("invalid filename")
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%sEnter new filename (default: %s): %s", Y, oldFilename, N)
	newFilename, _ := reader.ReadString('\n')
	newFilename = strings.TrimSpace(newFilename)
	return renameFile(oldFilename, newFilename, oldFilename)
}

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}
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
