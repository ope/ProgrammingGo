// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "Stdin", fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("\t%s\n", fileNames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName string, fileNames map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		if counts[text] == 1 {
			fileNames[text] = fileName
		} else {
			if !strings.Contains(fileNames[text], fileName) {
				fileNames[text] += " " + fileName
			}
		}
		// NOTE: ignoring potential errors from input.Err()
	}
}
