package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var errUsage = fmt.Errorf("Usage: %s <file1> <file2> ... <fileN>", os.Args[0])

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return errUsage
	}

	paths := os.Args[1:]
	files := make([]io.Reader, len(paths))
	for i, path := range paths {
		var err error
		files[i], err = os.Open(path)
		if err != nil {
			return fmt.Errorf("Error reading file %s: %s", path, err)
		}
	}

	for _, line := range Inter(files) {
		fmt.Println(line)
	}

	return nil
}

func Inter(readers []io.Reader) []string {
	count := make(map[string]int)

	for _, reader := range readers {
		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			count[line]++
		}
	}

	result := make([]string, 0, len(count))
	for line := range count {
		if count[line] != len(readers) {
			continue
		}
		result = append(result, line)
	}
	sort.Strings(result)

	return result
}
