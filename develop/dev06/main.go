package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fields := flag.String("f", "", "выбрать колонки")
	delimiter := flag.String("d", `\t`, "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// Проверка наличия флага -f
	if *fields == "" {
		fmt.Println("Необходимо указать флаг -f для выбора")
		os.Exit(1)
	}

	// Чтение строк из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка наличия разделителя в строке
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}
		// Разбиение строки на колонки
		columns := strings.Split(line, *delimiter)

		// Выбор указанных полей (колонок)
		selectedColumns := []string{}
		for _, field := range strings.Split(*fields, ",") {
			index := parseFieldIndex(field)

			if index > 0 && index <= len(columns) {
				selectedColumns = append(selectedColumns, columns[index-1])
			}
		}

		// Вывод выбранных колонок
		fmt.Println(selectedColumns)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error while reading STDIN:", err)
		os.Exit(1)
	}
}

// Функция для преобразования индекса поля (колонки) в число
func parseFieldIndex(field string) int {
	index := 0
	fmt.Sscanf(field, "%d", &index)
	return index
}
