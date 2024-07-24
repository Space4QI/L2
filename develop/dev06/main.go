package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Структура для хранения опций cut
type cutOptions struct {
	fields    string
	delimiter string
	separated bool
}

// Функция для парсинга флагов командной строки и создания структуры cutOptions
func parseFlags() cutOptions {
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	return cutOptions{
		fields:    *fields,
		delimiter: *delimiter,
		separated: *separated,
	}
}

// Функция для выполнения разрезания строк на основе опций cut
func cut(options cutOptions, input []string) []string {
	var result []string

	fields := parseFields(options.fields)
	for _, line := range input {
		if options.separated && !strings.Contains(line, options.delimiter) {
			continue
		}

		columns := strings.Split(line, options.delimiter)
		if len(fields) == 0 {
			result = append(result, line)
			continue
		}

		var selectedColumns []string
		for _, field := range fields {
			if field-1 < len(columns) {
				selectedColumns = append(selectedColumns, columns[field-1])
			}
		}
		result = append(result, strings.Join(selectedColumns, options.delimiter))
	}

	return result
}

// Функция для парсинга полей в виде строки в срез целых чисел
func parseFields(fields string) []int {
	var result []int
	if fields == "" {
		return result
	}

	for _, field := range strings.Split(fields, ",") {
		var fieldIndex int
		fmt.Sscanf(field, "%d", &fieldIndex)
		result = append(result, fieldIndex)
	}
	return result
}

// Основная функция, которая считывает входные данные и вызывает cut
func main() {
	options := parseFlags()

	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	result := cut(options, input)
	for _, line := range result {
		fmt.Println(line)
	}
}
