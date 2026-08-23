package fileutil

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
)

// ReadFile reads the entire content of a file and returns it as a string
func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadLines reads a file line by line and returns the lines as a slice of strings
func ReadLines(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists checks if a directory exists
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// ListFiles lists all files in a directory
func ListFiles(dirname string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, filepath.Join(dirname, entry.Name()))
		}
	}

	return files, nil
}

// FindFiles finds files matching a pattern in a directory
func FindFiles(dirname, pattern string) ([]string, error) {
	matches, err := filepath.Glob(filepath.Join(dirname, pattern))
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// ReadChunks reads a file in chunks of specified size
func ReadChunks(filename string, chunkSize int) (chunks [][]byte, err error) {
	if chunkSize <= 0 {
		return nil, errors.New("chunk size must be positive")
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); err == nil {
			err = closeErr
		}
	}()

	buffer := make([]byte, chunkSize)

	for {
		bytesRead, readErr := file.Read(buffer)
		if readErr != nil && readErr != io.EOF {
			return nil, readErr
		}

		if bytesRead == 0 {
			break
		}

		// Create a copy of the buffer to avoid overwriting it in the next iteration
		chunk := make([]byte, bytesRead)
		copy(chunk, buffer[:bytesRead])
		chunks = append(chunks, chunk)

		if readErr == io.EOF {
			break
		}
	}

	return chunks, nil
}
