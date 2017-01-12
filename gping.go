package main

// Package go-sh provides shell commands from golang

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/tatsushid/go-fastping"
)

// Server struct defines what is needed to ping servers
type Server struct {
	name     string
	pingable bool
}

func (s *Server) isPingable() bool {
	p := fastping.NewPinger()
	var status bool
	//name := "127.0.0.1"
	name := s.name
	ra, err := net.ResolveIPAddr("ip:icmp", name)
	if err != nil {
		return false
	}
	p.AddIPAddr(ra)
	//p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		//fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		//fmt.Printf("%s: is pingable\n", name)
		status = true
	}
	p.OnIdle = func() {
		//	fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
		status = false
	}
	return status
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
		if server.pingable == true {
			fmt.Printf("%s: is pingable\n", server.name)
		} else {
			fmt.Printf("%s: is NOT pingable\n", server.name)
		}
	}
}
