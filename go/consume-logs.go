package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// adapted from https://www.linkedin.com/pulse/golang-simple-tail-tagir-ismagilov/
func main() {
	filename := "../shared-fs/user-logs.log"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file: ", filename)
		return
	}
	defer file.Close()
	offset, err := file.Seek(0, io.SeekEnd)
	buffer := make([]byte, 1024, 1024)
	for {
		readBytes, err := file.ReadAt(buffer, offset)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading lines:", err)
				break
			}
		}
		offset += int64(readBytes)
		if readBytes != 0 {
			s := string(buffer[:readBytes])
			fmt.Printf("%s", s)
		}
		time.Sleep(time.Second)
	}
}
