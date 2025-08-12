package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileError struct {
	filename string
	line     int
	err      string
}

func (fe FileError) Error() string {
	return fmt.Sprintf("error processing line: %d, for file: %s, %s",
		fe.line, fe.filename, fe.err)
}

func processFile(name string) ([]int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s, %w", name, err)
	}

	nums := make([]int, 0)
	lineCt := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCt++
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, FileError{
				filename: name,
				line:     lineCt,
				err:      fmt.Sprintf("%s", err),
			}
		}

		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file %s, %w", name, err)
	}

	return nums, nil
}

func safeProcessFile(filename string) {
	fmt.Printf("📁 Processing file: %s\n", filename)

	numbers, err := processFile(filename)
	if err != nil {
		// Different handling based on error type
		if procErr, ok := err.(FileError); ok {
			fmt.Printf("❌ Data Error: %v\n", procErr)
			fmt.Println("💡 Suggestion: Check the file format at line", procErr.line)
		} else {
			fmt.Printf("❌ System Error: %v\n", err)
			fmt.Println("💡 Suggestion: Check if file exists and you have permissions")
		}
		return
	}

	// Success! Process the numbers
	fmt.Printf("✅ Successfully processed %d numbers: %v\n", len(numbers), numbers)
}

func main() {
	// This demonstrates graceful error handling
	testFiles := []string{"numbers.txt", "nonexistent.txt", "invalid_data.txt"}

	for _, filename := range testFiles {
		safeProcessFile(filename)
		fmt.Println(strings.Repeat("-", 50))
	}
}
