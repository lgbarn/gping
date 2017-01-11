package main

// Package go-sh provides shell commands from golang

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Server struct defines what is needed to ping servers
type Server struct {
	name  string
	ip    string
	valid bool
}

func checkValidIP(ip string) bool {
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		return false
	}
	return true
}

func lookupIP(name string) string {
	return name
}

func lookupname(ip string) string {
	return ip
}

//func (p *Server) populate(ip string) bool {
//trial := net.ParseIP(ip)
//if trial.To4() == nil {
//fmt.Printf("%v is not an IPv4 address\n", trial)
//return false
//}
//return false
//}

// main This program is used to run "ps -ef | grep <process>"
func main() {
	var stdInStat bool
	var servers []*Server
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdInStat = true
		stdInScanner := bufio.NewScanner(os.Stdin)
		for stdInScanner.Scan() {
			//fmt.Println(stdInScanner.Text())
			if err := stdInScanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}
	}
	var path string
	if os.Args != nil && len(os.Args) > 1 {
		path = os.Args[1]
		//path := "testfile"
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var line string
		for scanner.Scan() {
			line = scanner.Text()
			//var item string
			if checkValidIP(line) {
				//fmt.Printf("%s:valid\n", line)
				s := Server{ip: line}
				servers = append(servers, &s)
			} else {
				//fmt.Printf("%s:NOTvalid\n", line)
				s := Server{name: line}
				servers = append(servers, &s)
			}
		}
	}
	if os.Args != nil && len(os.Args) == 1 && stdInStat == false {
		fmt.Fprintln(os.Stderr, "No input provided")
	}
	for index, server := range servers {
		fmt.Printf("%d, %s, %s, %v\n", index, server.name, server.ip, server.valid)
	}
	//for server := range servers {
	//fmt.Printf("%T: %vs\n", server, server)
	//}
}
