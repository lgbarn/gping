package main

// Package go-sh provides shell commands from golang

import (
	"bufio"
	"fmt"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// main This program is used to run "ps -ef | grep <process>"
func main() {
	file := "testfile"
	contents, err := readLines(file)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%T : %v\n", contents, contents)
	for line := range contents {
		fmt.Printf("%T : %v\n", contents[line], contents[line])
	}
}
