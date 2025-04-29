package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func OpenFile(path string) io.ReadCloser {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func BytesCount(file io.ReadCloser) int {
	byte := make([]byte, 1)
	bytes := 0
	for {
		_, err := file.Read(byte)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		bytes++
	}
	return bytes
}

func LinesCount(file io.ReadCloser) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			log.Fatal(err)
		}
	}
}

func CountWords(file io.ReadCloser) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	count := 0

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return count
}
