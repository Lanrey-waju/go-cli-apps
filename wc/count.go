package main

import (
	"bufio"
	"io"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	switch {
	case countBytes:
		scanner.Split(bufio.ScanBytes)
	case !countLines:
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
