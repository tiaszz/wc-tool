// This package gets the user flag from the command line
// These are the flags for the command line:
// -c to count the number of bytes
// -l to count the number of lines
// -w to count the number of words
// -m to count the number of characters
package cli

import (
	"flag"
	"os"
)

type Cli struct {
	CheckBytes bool
	CheckLines bool
	CheckWords bool
	CheckChars bool
	FileName   string
}

func NewCli() *Cli {
	checkBytes := flag.Bool("c", false, "count the number of bytes in a file")
	checkLines := flag.Bool("l", false, "count the number of lines in a file")
	checkWords := flag.Bool("w", false, "count the number of words in a file")
	checkChars := flag.Bool("m", false, "count the number of characters in a file")

	flag.Parse()

	return &Cli{
		CheckBytes: *checkBytes,
		CheckLines: *checkLines,
		CheckWords: *checkWords,
		CheckChars: *checkChars,
		FileName:   determineFileName(),
	}
}

func HasProvided() bool {
	return len(os.Args) > 1
}

func determineFileName() string {
	if flag.NArg() > 0 {
		return flag.Arg(0)
	}
	return ""
}
