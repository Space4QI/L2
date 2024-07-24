package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Структура для хранения опций grep
type grepOptions struct {
	after      int
	before     int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
}

// Функция для парсинга флагов командной строки и создания структуры grepOptions
func parseFlags() grepOptions {
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")

	flag.Parse()

	pattern := flag.Arg(0)

	// Если установлен флаг -C, устанавливаем значения для -A и -B
	if *context > 0 {
		if *after == 0 {
			after = context
		}
		if *before == 0 {
			before = context
		}
	}

	return grepOptions{
		after:      *after,
		before:     *before,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
		pattern:    pattern,
	}
}

// Функция для выполнения фильтрации строк на основе опций grep
func grep(options grepOptions, input []string) []string {
	var result []string
	var count int
	var matchedLines []int

	// Игнорирование регистра
	if options.ignoreCase {
		options.pattern = strings.ToLower(options.pattern)
	}

	for i, line := range input {
		match := false
		if options.ignoreCase {
			line = strings.ToLower(line)
		}

		// Точное совпадение
		if options.fixed {
			match = (line == options.pattern)
		} else {
			match = strings.Contains(line, options.pattern)
		}

		// Инверсия совпадения
		if options.invert {
			match = !match
		}

		if match {
			matchedLines = append(matchedLines, i)
			if !options.count {
				if options.lineNum {
					result = append(result, fmt.Sprintf("%d:%s", i+1, input[i]))
				} else {
					result = append(result, input[i])
				}
			}
			count++
		}
	}

	// Возвращение количества совпадений
	if options.count {
		result = append(result, fmt.Sprintf("%d", count))
		return result
	}

	// Печать строк до и после совпадения
	if options.before > 0 || options.after > 0 {
		var finalResult []string
		seen := make(map[int]bool)

		for _, lineIdx := range matchedLines {
			start := lineIdx - options.before
			if start < 0 {
				start = 0
			}
			end := lineIdx + options.after
			if end >= len(input) {
				end = len(input) - 1
			}

			for j := start; j <= end; j++ {
				if !seen[j] {
					if options.lineNum {
						finalResult = append(finalResult, fmt.Sprintf("%d:%s", j+1, input[j]))
					} else {
						finalResult = append(finalResult, input[j])
					}
					seen[j] = true
				}
			}
		}
		return finalResult
	}

	return result
}

// Основная функция, которая считывает входные данные и вызывает grep
func main() {
	options := parseFlags()

	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	result := grep(options, input)
	for _, line := range result {
		fmt.Println(line)
	}
}
