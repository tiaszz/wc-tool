package main

import (
	"flag"
	"fmt"
)

func main() {
	numberBytes := flag.Bool("c", false, "Get the number of bytes of a file")
	numberLines := flag.Bool("l", false, "Get the number of lines of a file")
	numberWords := flag.Bool("w", false, "Get the number of words of a file")

	flag.Parse()
	fmt.Printf("%v\n%v\n%v\n", *numberBytes, *numberLines, *numberWords)
}
