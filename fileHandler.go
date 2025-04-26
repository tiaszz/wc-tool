package main

import (
	"io"
	"log"
	"os"
)

func OpenFile(path string) io.Reader {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	return f
}
