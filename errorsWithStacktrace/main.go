package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// A function that simulates an error
func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer file.Close()

	// Simulate reading file contents
	content := "file content here"
	return content, nil
}

// Another function that calls readFile and adds more context
func processFile(filename string) (string, error) {
	content, err := readFile(filename)
	if err != nil {
		return "", errors.Wrap(err, "failed to process file")
	}
	return content, nil
}

func main() {
	_, err := processFile("nonexistent.txt")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
