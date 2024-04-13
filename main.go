package main

import (
	"basic-wc/wc"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const usage = `Usage of wc:
  -c --byteCount      count of bytes in file
  -l --lineCount      count of lines in file
  -w --wordCount      count of words in file
  -m --characterCount count of character in file
`

func main() {
	var byteCount bool
	flag.BoolVar(&byteCount, "c", false, "count of bytes in file")
	flag.BoolVar(&byteCount, "byteCount", false, "count of bytes in file")

	var lineCount bool
	flag.BoolVar(&lineCount, "l", false, "count of lines in file")
	flag.BoolVar(&lineCount, "lineCount", false, "count of lines in file")

	var wordCount bool
	flag.BoolVar(&wordCount, "w", false, "count of words in file")
	flag.BoolVar(&wordCount, "wordCount", false, "count of words in file")

	var charCount bool
	flag.BoolVar(&charCount, "m", false, "count of chars in file")
	flag.BoolVar(&charCount, "charCount", false, "count of chars in file")

	flag.Usage = func() {
		fmt.Print(usage)
	}

	flag.Parse()

	filename := os.Args[len(os.Args)-1]
	filename = strings.TrimSpace(filename)

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Fprintf(os.Stderr, "Could not read file %s\n%s", filename, usage)
		os.Exit(1)
	}

	statTypes := wc.StatTypes{
		ByteRetrieve: byteCount,
		LineRetrieve: lineCount,
		WordRetrieve: wordCount,
		CharRetrieve: charCount,
	}

	if !statTypes.ByteRetrieve && !statTypes.LineRetrieve && !statTypes.WordRetrieve && !statTypes.CharRetrieve {
		statTypes.ByteRetrieve = true
		statTypes.LineRetrieve = true
		statTypes.WordRetrieve = true
	}

	statCounts, err := wc.GetStats(filename, statTypes)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Something went wrong: %v\n%s", err, usage)
		os.Exit(1)
	}

	fmt.Println(statCounts)
}
