package filemanager

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListFiles(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files
	testFiles := []string{"file1.txt", "file2.txt"}
	for _, file := range testFiles {
		filePath := filepath.Join(tempDir, file)
		_, err := os.Create(filePath)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	err = ListFiles(tempDir)
	if err != nil {
		t.Errorf("Listfiles failed: %v", err)
	}
}

func TestRenameFile(t *testing.T) {
	// Create test file
	tempDir, err := os.MkdirTemp("", "test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "file1.txt")
	newFilePath := filepath.Join(tempDir, "file2.txt")
	testfile, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer testfile.Close()
	testfile.Close()

	// Capture info for inode to ensure file hasn't changed after the rename
	fileInfoBefore, err := os.Stat(filePath)
	if err != nil {
		t.Fatalf("Failed to stat original file: %v", err)
	}
	// RenameFile failed
	// time.Sleep(100 * time.Millisecond)
	err = RenameFile(filePath, newFilePath)
	if err != nil {
		t.Errorf("RenameFile failed: %v", err)
		return
	}
	// Capture info for inode comparison
	fileInfoAfter, err := os.Stat(newFilePath)
	if err != nil {
		t.Fatalf("Failed to stat renamed file: %v", err)
	}

	if os.SameFile(fileInfoBefore,fileInfoAfter) {
		t.Errorf("file inode changed after rename")
	}
}
