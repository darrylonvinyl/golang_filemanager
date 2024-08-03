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
			args := os.Args[2:]
			if len(args) > 1 {
				fmt.Println("cd: too many arguments")
				continue
			}
			dir := "."
			if len(args) == 1 {
				dir = args[0]
			}
			err := filemanager.ChangeDirectory(dir)
			if err != nil {
				fmt.Println("Error changing directory:", err)
			} else {
				cwd, _ = os.Getwd()
			}
		}
	}

}
