package main

import (
	"bufio"
	"os"
	"log"
)

func ReadLines(filename string) []string {
	var fileLines []string
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed opening file: %s\n", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	return fileLines
}
