package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("failed to create listener")
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("failed to establish connection")
		}
		log.Print("connection established...")
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		conn.Close()
		fmt.Println("connection closed")
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	strChan := make(chan string)
	var currentLine string
	go func() {
		defer f.Close()
		defer close(strChan)
		for {
			buffer := make([]byte, 8)
			var bytesRead int
			var err error
			if bytesRead, err = f.Read(buffer); err == io.EOF {
				if currentLine != "" {
					strChan <- currentLine
				}
				break
			} else if err != nil {
				log.Fatal("error reading bytes")
			}
			byteString := string(buffer[:bytesRead])
			parts := strings.Split(byteString, "\n")
			for i := range len(parts) - 1 {
				currentLine += parts[i]
				strChan <- currentLine
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	}()
	return strChan
}
