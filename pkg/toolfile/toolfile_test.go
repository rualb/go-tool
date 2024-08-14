package toolfile

import (
	"os"
	"testing"
)

func TestGetTempFile(t *testing.T) {
	filename, err := GetTempFile("test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename) // Clean up

	if !Exists(filename) {
		t.Fatalf("Temp file %s does not exist", filename)
	}
}

func TestDelete(t *testing.T) {
	filename, err := GetTempFile("test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename) // Clean up

	err = Delete(filename)
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	if Exists(filename) {
		t.Fatalf("File %s still exists after deletion", filename)
	}
}

func TestRename(t *testing.T) {
	oldPath, err := GetTempFile("testrename")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(oldPath) // Clean up after test

	newPath := oldPath + "-renamed"

	if err := Rename(oldPath, newPath); err != nil {
		t.Fatalf("Rename failed: %v", err)
	}

	if !Exists(newPath) {
		t.Fatalf("File was not renamed to %s", newPath)
	}

	if Exists(oldPath) {
		t.Fatalf("Old file %s still exists", oldPath)
	}

	// Clean up
	os.Remove(newPath)
}

func TestExists(t *testing.T) {
	filename, err := GetTempFile("test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename) // Clean up

	if !Exists(filename) {
		t.Fatalf("File %s should exist", filename)
	}

	if Exists("nonexistent-file.tmp") {
		t.Fatalf("Nonexistent file should not exist")
	}
}

func TestWriteAllText(t *testing.T) {
	filename, err := GetTempFile("test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename) // Clean up

	text := "Hello, World!"
	err = WriteAllText(filename, text)
	if err != nil {
		t.Fatalf("Failed to write text: %v", err)
	}

	readText, err := ReadAllText(filename)
	if err != nil {
		t.Fatalf("Failed to read text: %v", err)
	}

	if readText != text {
		t.Fatalf("Expected %q but got %q", text, readText)
	}
}

func TestWriteBytes(t *testing.T) {
	filename, err := GetTempFile("test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename) // Clean up

	data := []byte{1, 2, 3, 4, 5}
	err = WriteBytes(filename, data)
	if err != nil {
		t.Fatalf("Failed to write bytes: %v", err)
	}

	readData, err := ReadAllBytes(filename)
	if err != nil {
		t.Fatalf("Failed to read bytes: %v", err)
	}

	if string(readData) != string(data) {
		t.Fatalf("Expected %v but got %v", data, readData)
	}
}

func TestAppendText(t *testing.T) {
	filename, err := GetTempFile("test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename) // Clean up

	initialText := "Hello"
	err = WriteAllText(filename, initialText)
	if err != nil {
		t.Fatalf("Failed to write initial text: %v", err)
	}

	appendText := " World!"
	err = AppendText(filename, appendText)
	if err != nil {
		t.Fatalf("Failed to append text: %v", err)
	}

	expectedText := initialText + appendText
	readText, err := ReadAllText(filename)
	if err != nil {
		t.Fatalf("Failed to read text: %v", err)
	}

	if readText != expectedText {
		t.Fatalf("Expected %q but got %q", expectedText, readText)
	}
}

func TestReadAllLines(t *testing.T) {
	filename, err := GetTempFile("test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename) // Clean up

	lines := []string{"Line 1", "Line 2", "Line 3"}
	for _, line := range lines {
		err = AppendText(filename, line+"\n")
		if err != nil {
			t.Fatalf("Failed to append line: %v", err)
		}
	}

	readLines, err := ReadAllLines(filename)
	if err != nil {
		t.Fatalf("Failed to read lines: %v", err)
	}

	if len(readLines) != len(lines) {
		t.Fatalf("Expected %d lines but got %d", len(lines), len(readLines))
	}

	for i, line := range lines {
		if readLines[i] != line {
			t.Fatalf("Expected line %d to be %q but got %q", i+1, line, readLines[i])
		}
	}
}
