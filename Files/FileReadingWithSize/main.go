package main

// This program will read a text file.
// The difference between this directory and the ReadFile directory
// is that this program will utilize the FileInfo type to get the size of the
// file and use that to set the length of the slice rather than just hardcoding a value and
// be unsure of the actual size of the file.  You could either waste space, or you could
// end up not reading the whole file.

// Wonder if there is a better way to do this (._.)

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	ErrorCheck(err)

	fileInfo, err := file.Stat()
	ErrorCheck(err)

	data := CreateSlice(fileInfo.Size())

	ReadFile(file, data)

	Close(file)

	fmt.Println("All done!")
}

func ReadFile(file *os.File, data []byte) {
	bytesRead, err := file.Read(data)
	ErrorCheck(err)

	DisplayData(data, bytesRead)
}

func DisplayData(data []byte, bytesRead int) {
	fmt.Printf("Bytes Read: %d\n%s\n", bytesRead, data)
}

func Close(file *os.File) {
	err := file.Close()
	ErrorCheck(err)
}

func CreateSlice(size int64) []byte {
	data := make([]byte, size)

	return data
}

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
