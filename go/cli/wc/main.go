package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count lines")
	bytes := flag.Bool("b", false, "count bytes")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	f := splitFunc(countLines, countBytes)
	scanner.Split(f)
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}

func splitFunc(countLines, countBytes bool) bufio.SplitFunc {
	switch {
	case countBytes:
		return bufio.ScanBytes
	case countLines:
		return bufio.ScanLines
	default:
		return bufio.ScanWords
	}
}
