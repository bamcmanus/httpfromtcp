package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer file.Close()

	for line := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", line)
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
			for i:=range len(parts)-1 {
				currentLine += parts[i]
				strChan <- currentLine
				currentLine = ""
			}
			currentLine += parts[len(parts)-1]
		}
	}()
	return strChan
}
