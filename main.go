package main

import (
	"cli_filemanager/filemanager"
	"fmt"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory: ", err)
	}

	for {
		fmt.Printf("Current directory: %s\n", cwd)
		fmt.Print("> ")

		// Read user input
		var command string
		fmt.Scanln(&command)

		// Handle commands (e.g.,  "ls", "cd")
		switch command {
		case "ls":
			err := filemanager.ListFiles(cwd)
			if err != nil {
				fmt.Println("Error listing files:", err)
			}
		case "cd":
			if len(os.Args) > 2 {
				dir := os.Args[2]
				err := filemanager.ChangeDirectory(dir)
				if err != nil {
					fmt.Println("Error changing directory:",err)
				} else {
					cwd, err = os.Getwd()
					if err != nil {
						fmt.Printf("Error getting working directory: %v", err)
					}
				}
			} else {
				fmt.Println("cd: missing directory argument")
			}
	case "pwd":
		filemanager.PrintWorkingDirectory()
	}

	}
}
