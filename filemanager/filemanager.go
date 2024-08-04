package filemanager

import (
	"fmt"
	"os"
)

func ListFiles(dirPath string) error {
	// Open the directory
	d, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer d.Close()

	// Read the directory entries
	fileInfos, err := d.ReadDir(0)
	if err != nil {
		return err
	}

	// Print file and directory names
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}

	return nil
}

func CreateFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func RenameFile(oldFilePath, newFilePath string) error {
	return os.Rename(oldFilePath,newFilePath)
}

func ChangeDirectory(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}
	return nil
}

func PrintWorkingDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}
	fmt.Println(cwd)
	return cwd, nil
}
