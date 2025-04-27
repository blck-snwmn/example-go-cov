package fileutil

import (
	"bufio"
	"os"
	"path/filepath"
)

// WriteFile writes a string to a file
func WriteFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// WriteLines writes a slice of strings to a file, one line per string
func WriteLines(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// AppendToFile appends a string to a file
func AppendToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// AppendLines appends a slice of strings to a file, one line per string
func AppendLines(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// CreateDir creates a directory and all parent directories if they don't exist
func CreateDir(dirname string) error {
	return os.MkdirAll(dirname, 0755)
}

// CopyFile copies a file from src to dst
func CopyFile(src, dst string) error {
	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	if err := CreateDir(dstDir); err != nil {
		return err
	}

	// Read the source file
	content, err := ReadFile(src)
	if err != nil {
		return err
	}

	// Write the content to the destination file
	return WriteFile(dst, content)
}

// DeleteFile deletes a file
func DeleteFile(filename string) error {
	return os.Remove(filename)
}

// TruncateFile truncates a file to zero size
func TruncateFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
