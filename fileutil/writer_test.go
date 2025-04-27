package fileutil

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	// Create a temporary directory for our test files
	tempDir, err := os.MkdirTemp("", "writefile_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test file path
	testFile := tempDir + "/test_write.txt"
	testContent := "Hello, World!"

	// Test writing
	err = WriteFile(testFile, testContent)
	if err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	// Read back and verify content
	content, err := ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if content != testContent {
		t.Errorf("WriteFile() content = %q, want %q", content, testContent)
	}
}

func TestWriteLines(t *testing.T) {
	// Create a temporary directory for our test files
	tempDir, err := os.MkdirTemp("", "writelines_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test file path
	testFile := tempDir + "/test_write_lines.txt"
	testLines := []string{"Line 1", "Line 2", "Line 3"}
	expectedContent := "Line 1\nLine 2\nLine 3\n"

	// Test writing
	err = WriteLines(testFile, testLines)
	if err != nil {
		t.Fatalf("WriteLines() error = %v", err)
	}

	// Read back and verify content
	content, err := ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if content != expectedContent {
		t.Errorf("WriteLines() content = %q, want %q", content, expectedContent)
	}
}

func TestAppendToFile(t *testing.T) {
	// Create a temporary directory for our test files
	tempDir, err := os.MkdirTemp("", "appendfile_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test file path
	testFile := tempDir + "/test_append.txt"
	initialContent := "Initial content\n"
	appendedContent := "Appended content"
	expectedContent := initialContent + appendedContent

	// Create initial file
	err = WriteFile(testFile, initialContent)
	if err != nil {
		t.Fatalf("Failed to create initial test file: %v", err)
	}

	// Test appending
	err = AppendToFile(testFile, appendedContent)
	if err != nil {
		t.Fatalf("AppendToFile() error = %v", err)
	}

	// Read back and verify content
	content, err := ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	if content != expectedContent {
		t.Errorf("AppendToFile() content = %q, want %q", content, expectedContent)
	}
}

// We skip testing AppendLines, CreateDir, CopyFile, DeleteFile, and TruncateFile
// to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
