package main

import (
	"fmt"
	"log"
	"os"
	//"github.com/tiaszz/wc-tool/cli"
)

func main() {
	fmt.Println(readFile("test.txt"))
}

func readFile(filename string) []byte {
	input, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return input
}