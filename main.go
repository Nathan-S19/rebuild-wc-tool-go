package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	linesFlag bool
	wordsFlag bool
	bytesFlag bool
	charsFlag bool
)

func init() {
	flag.BoolVar(&linesFlag, "l", false, "Print the newline counts")
	flag.BoolVar(&linesFlag, "lines", false, "Print the newline counts")
	flag.BoolVar(&wordsFlag, "w", false, "Print the word counts")
	flag.BoolVar(&wordsFlag, "words", false, "Print the word counts")
	flag.BoolVar(&bytesFlag, "c", false, "Print the byte counts")
	flag.BoolVar(&bytesFlag, "bytes", false, "Print the byte counts")
	flag.BoolVar(&charsFlag, "m", false, "Print the character counts")
	flag.BoolVar(&charsFlag, "chars", false, "Print the character counts")
}

func countBytes(content string) int {
	return len(content)
}

func countLines(content string) int {
	return len(strings.Split(content, "\n")) - 1
}

func countWords(content string) int {
	return len(strings.Fields(content))
}

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Please provide a filename.")
		return
	}
	filename := flag.Args()[0]

	content, err := ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	byteCount := countBytes(content)
	lineCount := countLines(content)
	wordCount := countWords(content)

	// Print results based on flags
	if linesFlag {
		fmt.Printf("Lines: %d\n", lineCount)
	}
	if wordsFlag {
		fmt.Printf("Words: %d\n", wordCount)
	}
	if bytesFlag {
		fmt.Printf("Bytes: %d\n", byteCount)
	}
	if charsFlag {
		fmt.Printf("Characters: %d\n", byteCount)
	}

	if !linesFlag && !wordsFlag && !bytesFlag && !charsFlag {
		fmt.Printf("Lines: %d  Words: %d  Bytes: %d\n", lineCount, wordCount, byteCount)
	}
}
