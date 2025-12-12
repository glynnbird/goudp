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
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		return "", err
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()
	message := []byte(fmt.Sprintf("%v\n", messageStr))
	_, err = conn.Write(message)
	if err != nil {
		return "", err
	}

	// wait up to 2ms for a reply
	conn.SetReadDeadline(time.Now().Add(2 * time.Millisecond))
	buffer := make([]byte, 8)
	_, _, err = conn.ReadFromUDP(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), err
}
