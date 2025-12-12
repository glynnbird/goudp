package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: goudp <message>")
		os.Exit(1)
	}
	reply, err := sendAndReceive(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(reply)
}

func sendAndReceive(messageStr string) (string, error) {
	// send message via UDP
	serverAddr := os.Getenv("UDP_SERVER")
	if serverAddr == "" {
		serverAddr = "127.0.0.1:8080"
	}
	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	message := []byte(messageStr + "\n")
	_, err = conn.Write(message)
	if err != nil {
		return "", err
	}

	// wait up to 2ms for a reply
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}
