package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Tes5t"
	fmt.Println(unpackString(str))
}

func unpackString(s string) string {
	var result strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			return ""
		}

		if i+1 < len(runes) && runes[i+1] >= '0' && runes[i+1] <= '9' {
			count, _ := strconv.Atoi(string(runes[i+1]))
			result.WriteString(strings.Repeat(string(runes[i]), count))
			i++
		} else {
			result.WriteRune(runes[i])
		}
	}

	return result.String()
}
