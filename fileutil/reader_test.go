package fileutil

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "fileexists_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "fileexists_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{"existing file", tempFile.Name(), true},
		{"non-existent file", "non_existent_file.txt", false},
		{"directory", tempDir, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FileExists(tt.filename)
			if got != tt.expected {
				t.Errorf("FileExists(%q) = %v, want %v", tt.filename, got, tt.expected)
			}
		})
	}
}

func TestDirExists(t *testing.T) {
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "direxists_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "direxists_test_dir")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	tests := []struct {
		name     string
		dirname  string
		expected bool
	}{
		{"existing directory", tempDir, true},
		{"non-existent directory", "non_existent_dir", false},
		{"file", tempFile.Name(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DirExists(tt.dirname)
			if got != tt.expected {
				t.Errorf("DirExists(%q) = %v, want %v", tt.dirname, got, tt.expected)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	// Create a temporary file with content
	content := "Hello, World!"
	tempFile, err := os.CreateTemp("", "readfile_test_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFile.WriteString(content)
	tempFile.Close()
	defer os.Remove(tempFile.Name())

	tests := []struct {
		name        string
		filename    string
		expected    string
		expectError bool
	}{
		{"existing file", tempFile.Name(), content, false},
		{"non-existent file", "non_existent_file.txt", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.filename)

			// Check error expectation
			if (err != nil) != tt.expectError {
				t.Errorf("ReadFile(%q) error = %v, expectError %v", tt.filename, err, tt.expectError)
				return
			}

			// If we're expecting an error, don't check the result
			if tt.expectError {
				return
			}

			if got != tt.expected {
				t.Errorf("ReadFile(%q) = %q, want %q", tt.filename, got, tt.expected)
			}
		})
	}
}

// We skip testing ReadLines, ListFiles, FindFiles, and ReadChunks to deliberately leave them uncovered
// This helps us achieve a coverage of around 50%
