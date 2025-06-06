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

	var currentLine string
	for {
		buffer := make([]byte, 8)
		var bytesRead int
		if bytesRead, err = file.Read(buffer); err == io.EOF {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)	
			}
			break
		} else if err != nil {
			log.Fatal("error reading bytes")
		}
		byteString := string(buffer[:bytesRead])
		parts := strings.Split(byteString, "\n")
		for i:=0; i < len(parts)-1; i++ {
			currentLine += parts[i]
			fmt.Printf("read: %s\n", currentLine)	
			currentLine = ""
		}
		currentLine += parts[len(parts)-1]
	}
}
