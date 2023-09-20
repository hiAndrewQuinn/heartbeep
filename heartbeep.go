package main

import (
  "time"
  "os"
	"flag"
	"log"
	"net"
)

// isAlive checks if a given host and port are reachable.
func isAlive(host string, port string) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	ipAddr := flag.String("ip", "10.0.0.2", "IP Address of the machine to monitor")
	port := flag.String("port", "22", "Port to check on the machine")
	flag.Parse()

	if !isAlive(*ipAddr, *port) {
		log.Printf("Machine %s on port %s is down!\n", *ipAddr, *port)
    os.Exit(1)
	}
  os.Exit(0)
}

