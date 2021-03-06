package main

// Package go-sh provides shell commands from golang

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
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
	name := s.name
	ra, err := net.ResolveIPAddr("ip:icmp", name)
	if err != nil {
		return false
	}
	p.AddIPAddr(ra)
	if runtime.GOOS == "linux" {
		p.Network("ip")
	} else {
		p.Network("udp")
	}
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		status = true
	}
	p.OnIdle = func() {
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
		status = false
	}
	if status == true {
		fmt.Printf("%s: is pingable\n", name)
	} else {
		fmt.Printf("%s: is NOT pingable\n", name)
	}
	return status
}

// main This program is used to run "ps -ef | grep <process>"
func main() {
	var stdInStat bool
	var stdInServer string
	var path string
	var line string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdInStat = true
		stdInScanner := bufio.NewScanner(os.Stdin)
		stdInScanner.Split(bufio.ScanWords)
		for stdInScanner.Scan() {
			stdInServer = stdInScanner.Text()
			s := Server{name: stdInServer}
			s.pingable = s.isPingable()
			if err := stdInScanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		}
	}
	if os.Args != nil && len(os.Args) > 1 {
		path = os.Args[1]
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line = scanner.Text()
			s := Server{name: line}
			s.pingable = s.isPingable()
		}
	}
	if os.Args != nil && len(os.Args) == 1 && stdInStat == false {
		fmt.Fprintln(os.Stderr, "No input provided")
	}
}
