package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	column := flag.Int("k", 0, "column for sorting")
	numeric := flag.Bool("n", false, "sort by numeric value")
	reverse := flag.Bool("r", false, "sort in reverse order")
	unique := flag.Bool("u", false, "do not output duplicate lines")

	flag.Parse()

	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	lines, err := readLines(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	sort.Slice(lines, func(i, j int) bool {
		return compareLines(lines[i], lines[j], *column, *numeric) < 0
	})
	if *reverse {
		reverseLines(lines)
	}
	if *unique {
		lines = removeDuplicates(lines)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

// Чтение строк из ридера
func readLines(reader io.Reader) ([]string, error) {
	lines := []string{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Сравнение строк для сортировки
func compareLines(line1, line2 string, column int, numeric bool) int {
	fields1 := strings.Fields(line1)
	fields2 := strings.Fields(line2)

	if column > 0 && column <= len(fields1) && column <= len(fields2) {
		value1 := fields1[column-1]
		value2 := fields2[column-1]

		if numeric {
			num1, err := strconv.Atoi(value1)
			if err == nil {
				num2, err := strconv.Atoi(value2)
				if err == nil {
					if num1 < num2 {
						return -1
					} else if num1 > num2 {
						return 1
					} else {
						return 0
					}
				}
			}
		} else {
			if value1 < value2 {
				return -1
			} else if value1 > value2 {
				return 1
			} else {
				return 0
			}
		}
	}

	return strings.Compare(line1, line2)
}

// Реверс строки
func reverseLines(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

// Удаление повторяющихся строк
func removeDuplicates(lines []string) []string {
	result := []string{}
	seen := make(map[string]bool)
	for _, line := range lines {
		if !seen[line] {
			result = append(result, line)
			seen[line] = true
		}
	}
	return result
}
