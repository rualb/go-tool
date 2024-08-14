package toolfile

import (
	"bufio"
	"fmt"
	"os"
)

// CreateTemp creates a temporary file and returns its name.
func CreateTemp(prefix string) (string, error) {
	// Create a temporary file in the default directory
	tempFile, err := os.CreateTemp("", prefix+"-*.tmp")
	if err != nil {
		return "", err
	}
	// Ensure the file is deleted when no longer needed
	//defer os.Remove(tempFile.Name())
	defer tempFile.Close()
	// Return the file name
	return tempFile.Name(), nil
}

func Delete(filepath string) error {
	err := os.Remove(filepath)

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
func Exists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil && !info.IsDir()
}

func Rename(oldPath, newPath string) error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("failed to rename file from %s to %s: %w", oldPath, newPath, err)
	}
	return nil
}

// writes all text to a file.
func WriteAllText(filepath, text string) error {
	err := os.WriteFile(filepath, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}

// writes all bytes to a file.
func WriteBytes(filepath string, data []byte) error {
	err := os.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// appends text to a file.

func AppendText(filepath, text string) error {
	// Open the file for appending; create it if it doesn't exist
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
func ReadAllText(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// reads the entire file content as a byte array.
func ReadAllBytes(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// reads the file content line by line.

func ReadAllLines(filepath string) ([]string, error) {
	var lines []string
	file, err := os.Open(filepath)
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
