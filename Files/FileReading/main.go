package main

/*
 * Notice how all of the contents aren't read.
 * This shall be fixed in the next example: FileReadingWithSize
 */

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	ErrorCheck(err)

	ReadFileContents(file)
}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFileContents(file *os.File) {
	data := make([]byte, 100)

	bytesRead, err := file.Read(data)
	ErrorCheck(err)

	OutputData(data, bytesRead)
}

func OutputData(data []byte, bytes int) {
	fmt.Printf("Bytes Read: %d\n%s\n", bytes, data)
}
