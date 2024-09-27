package toolzip

import (
	"bytes"
	"testing"
)

// TestZipUnZip tests the Zip and UnZip functions with directory paths.
func TestZipUnZip(t *testing.T) {
	// Define the files to be zipped, including directory paths
	filesToZip := map[string][]byte{
		"dir1/file1.txt": []byte("This is the content of file1."),
		"dir2/file2.txt": []byte("This is the content of file2."),
	}

	// Create the zip archive
	zipData, err := Zip(filesToZip)
	if err != nil {
		t.Fatalf("Zip() error = %v", err)
	}

	// Unzip the archive
	unzippedFiles, err := UnZip(zipData, 0)
	if err != nil {
		t.Fatalf("UnZip() error = %v", err)
	}

	// Check that the number of files matches
	if len(unzippedFiles) != len(filesToZip) {
		t.Errorf("UnZip() mismatch: got %d files, want %d", len(unzippedFiles), len(filesToZip))
	}

	// Check the contents of each file
	for path, expectedContent := range filesToZip {
		content, ok := unzippedFiles[path]
		if !ok {
			t.Errorf("UnZip() missing file: %s", path)
			continue
		}
		if !bytes.Equal(content, expectedContent) {
			t.Errorf("UnZip() content mismatch for file %s: got %s, want %s", path, content, expectedContent)
		}
	}
}

// TestZipEmpty tests the Zip function with an empty map and ensures it creates an empty zip archive.
func TestZipEmpty(t *testing.T) {
	filesToZip := map[string][]byte{}

	zipData, err := Zip(filesToZip)
	if err != nil {
		t.Fatalf("Zip() error = %v", err)
	}

	// Unzip the archive to check if it is empty
	unzippedFiles, err := UnZip(zipData, 0)
	if err != nil {
		t.Fatalf("UnZip() error = %v", err)
	}

	if len(unzippedFiles) != 0 {
		t.Errorf("UnZip() expected 0 files, got %d", len(unzippedFiles))
	}
}

// TestZipInvalidData tests the UnZip function with invalid zip data.
func TestZipInvalidData(t *testing.T) {
	invalidData := []byte("invalid zip data")

	_, err := UnZip(invalidData, 0)
	if err == nil {
		t.Error("UnZip() expected error, got nil")
	}
}
