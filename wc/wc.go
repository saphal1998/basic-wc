package wc

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type StatTypes struct {
	ByteRetrieve bool
	LineRetrieve bool
	WordRetrieve bool
	CharRetrieve bool
}

type StatCounts struct {
	StatTypes
	filename  string
	ByteCount uint64
	LineCount uint64
	WordCount uint64
	CharCount uint64
}

func (sc *StatCounts) String() string {
	output := ""

	if sc.LineRetrieve {
		output += fmt.Sprintf("\tlineCount(%d)", sc.LineCount)
	}
	if sc.WordRetrieve {
		output += fmt.Sprintf("\twordCount(%d)", sc.WordCount)
	}
	if sc.ByteRetrieve {
		output += fmt.Sprintf("\tbyteCount(%d)", sc.ByteCount)
	}
	if sc.CharRetrieve {
		output += fmt.Sprintf("\tcharCount(%d)", sc.CharCount)
	}
	if len(sc.filename) > 0 {
		output += fmt.Sprintf("\t%s", sc.filename)
	}

	output += "\n"
	return output
}

func GetStats(filename string, statTypes StatTypes) (*StatCounts, error) {
	statCounts := StatCounts{
		StatTypes: statTypes,
		filename:  filename,
	}
	var f *os.File
	if len(filename) == 0 {
		f = os.Stdin
	} else {
		// This is added here, because if we use the `:=` syntax in os.Open, the file is closed as soon as the else scope ends, which is not what we want
		var err error
		f, err = os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer f.Close()
	}

	byteCount := make(chan uint64, 0)
	wordCount := make(chan uint64, 0)
	runeCount := make(chan uint64, 0)
	lineCount := make(chan uint64, 0)

	if statCounts.ByteRetrieve {
		go getScanCount(f, bufio.ScanBytes, byteCount)
	}
	if statCounts.WordRetrieve {
		go getScanCount(f, bufio.ScanWords, wordCount)
	}
	if statCounts.CharRetrieve {
		go getScanCount(f, bufio.ScanRunes, runeCount)
	}
	if statCounts.LineRetrieve {
		go getScanCount(f, bufio.ScanLines, lineCount)
	}

	if statCounts.ByteRetrieve {
		statCounts.ByteCount = <-byteCount
	}
	if statCounts.WordRetrieve {
		statCounts.WordCount = <-wordCount
	}
	if statCounts.CharRetrieve {
		statCounts.CharCount = <-runeCount
	}
	if statCounts.LineRetrieve {
		statCounts.LineCount = <-lineCount
	}
	return &statCounts, nil
}

func getScanCount(r io.Reader, splitFunc bufio.SplitFunc, outputChan chan<- uint64) {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFunc)
	count := 0

	for scanner.Scan() {
		count += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	outputChan <- uint64(count)
}
