package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "Print +N lines after each match")
	before := flag.Int("B", 0, "Print +N lines before each match")
	context := flag.Int("C", 0, "Print ±N lines around each match")
	count := flag.Bool("c", false, "Count the number of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert the match")
	fixed := flag.Bool("F", false, "Exact string match")
	lineNum := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	pattern := flag.Arg(0)

	if pattern == "" {
		fmt.Println("Pattern is required")
		return
	}

	file, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var output []string
	var lineCount int

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println("Error reading file:", err)
			return
		}

		if *ignoreCase || *fixed {
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}

		match := strings.Contains(line, pattern) || (*fixed && strings.EqualFold(line, pattern))

		if (*invert && !match) || (!*invert && match) {
			lineCount++

			if *count {
				continue
			}

			if *lineNum {
				line = fmt.Sprintf("%d: %s", lineCount, line)
			}

			output = append(output, line)

			if *after > 0 {
				for i := 0; i < *after; i++ {
					nextLine, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Println("Error reading file:", err)
						return
					}
					output = append(output, nextLine)
				}
			}

			if *context > 0 {
				linesBefore := *context
				linesAfter := *context

				if *before > 0 {
					linesBefore = *before
				}

				if *after > 0 {
					linesAfter = *after
				}

				for i := 0; i < linesBefore; i++ {
					prevLine, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Println("Error reading file:", err)
						return
					}
					output = append([]string{prevLine}, output...)
				}

				for i := 0; i < linesAfter; i++ {
					nextLine, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Println("Error reading file:", err)
						return
					}
					output = append(output, nextLine)
				}
			}
		}

		if err == io.EOF {
			break
		}
	}

	for _, line := range output {
		fmt.Print(line)
	}
}
