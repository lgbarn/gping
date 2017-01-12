package main

// Package go-sh provides shell commands from golang

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Server struct defines what is needed to ping servers
type Server struct {
	name     string
	pingable bool
}

func (s *Server) isPingable() bool {
	name := s.name
	fmt.Printf("%s is pingable\n", name)
	return true
}

// main This program is used to run "ps -ef | grep <process>"
func main() {
	var stdInStat bool
	var servers []*Server
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdInStat = true
		stdInScanner := bufio.NewScanner(os.Stdin)
		for stdInScanner.Scan() {
			if err := stdInScanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}
	}
	var path string
	if os.Args != nil && len(os.Args) > 1 {
		path = os.Args[1]
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			s := Server{name: line}
			s.pingable = s.isPingable()
			servers = append(servers, &s)
		}
	}
	if os.Args != nil && len(os.Args) == 1 && stdInStat == false {
		fmt.Fprintln(os.Stderr, "No input provided")
	}
	for _, server := range servers {
		fmt.Printf("%s, %v\n", server.name, server.pingable)
	}
}
