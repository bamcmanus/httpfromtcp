package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatal("failed to resolve UDP address")
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("could not establish connection")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error reading from stdin:", err)
			continue
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			log.Println("error while writing bytes:", err)
		}
	}
}
