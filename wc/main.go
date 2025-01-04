package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "count lines in input")
	bytes := flag.Bool("b", false, "count number of bytes")

	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
