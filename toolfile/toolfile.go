package toolfile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// CreateTemp creates a temporary file and returns its name.
func CreateTemp(prefix string) (string, error) {
	// Create a temporary file in the default directory
	tempFile, err := os.CreateTemp("", prefix+"-*.tmp")
	if err != nil {
		return "", err
	}
	// Ensure the file is deleted when no longer needed
	// defer os.Remove(tempFile.Name())
	defer tempFile.Close()
	// Return the file name
	return tempFile.Name(), nil
}

func Delete(path string) error {
	path = filepath.Clean(path) // do Abs before
	err := os.Remove(path)

	// if err == nil {
	// 	return nil
	// }
	// if os.IsPermission(err) || os.IsNotExist(err) {
	// 	return err
	// }

	if err != nil {
		return err
	}
	return nil
}

// checks if a file exists at the specified path.
func Exists(path string) bool {
	path = filepath.Clean(path) // do Abs before
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && !info.IsDir()
}

func Rename(oldPath, newPath string) error {
	oldPath = filepath.Clean(oldPath) // do Abs before
	newPath = filepath.Clean(newPath) // do Abs before
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("failed to rename file from %s to %s: %w", oldPath, newPath, err)
	}
	return nil
}

// writes all text to a file.
func WriteAllText(file, text string) error {
	err := os.WriteFile(file, []byte(text), 0600)
	if err != nil {
		return err
	}
	return nil
}

// writes all bytes to a file.
func WriteBytes(path string, data []byte) error {
	path = filepath.Clean(path) // do Abs before
	err := os.WriteFile(path, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

// appends text to a file.

func AppendText(path, text string) error {
	path = filepath.Clean(path)
	// Open the file for appending; create it if it doesn't exist
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

// reads the entire file content as a text string.
func ReadAllText(path string) (string, error) {
	path = filepath.Clean(path)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// reads the entire file content as a byte array.
func ReadAllBytes(path string) ([]byte, error) {
	path = filepath.Clean(path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// reads the file content line by line.

func ReadAllLines(path string) ([]string, error) {
	path = filepath.Clean(path)
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
