package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world AI PR")

	readFile("sample.txt")
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("file --> ", file)

	file.Close()
}
