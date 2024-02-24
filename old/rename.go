package main

import (
	"fmt"
	"os"
	//"path/filepath"
)

func rename(args []string) {
	switch args[0] {
	case "--help", "-h", "-?":
		fmt.Println("Usage: rename <filename>")
	case "--version", "-v":
		fmt.Println("rename v1.0")
		fmt.Println("by PhateValleyman")
		fmt.Println("Jonas.Ned@outlook.com")
	default:
		if len(args) < 1 || !fileExists(args[0]) {
			fmt.Println("Usage: rename <filename>")
			return
		}

		var newfilename string
		fmt.Printf("%s: ", args[0])
		fmt.Scanln(&newfilename)
		err := os.Rename(args[0], newfilename)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "complete" {
		// Shell completion logic goes here
		// Note: Go doesn't have a direct equivalent to "compgen"
	} else {
		rename(os.Args[1:])
	}
}
