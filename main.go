package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define flags
	linesFlag := flag.Bool("l", false, "Count lines")
	wordsFlag := flag.Bool("w", false, "Count words")
	bytesFlag := flag.Bool("c", false, "Count bytes")
	flag.Parse()

	// Get the input files from the arguments
	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Usage: go-wc [options] file...")
		return
	}

	// Process each file
	for _, file := range files {
		err := processFile(file, *linesFlag, *wordsFlag, *bytesFlag)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", file, err)
		}
	}
}

func processFile(file string, countLines, countWords, countBytes bool) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	var lines, words, bytes int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		words += len(strings.Fields(line))
		bytes += len(line) + 1 // Add 1 for the newline character
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	// Display results
	fmt.Printf("%s:\n", file)
	if countLines {
		fmt.Printf("Lines: %d\n", lines)
	}
	if countWords {
		fmt.Printf("Words: %d\n", words)
	}
	if countBytes {
		fmt.Printf("Bytes: %d\n", bytes)
	}
	return nil
}
