package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Конвертирует строку с суффиксом в числовое значение
func parseHumanReadableNumber(s string) (float64, error) {
	units := map[byte]float64{
		'K': 1 << 10, // Килобайт
		'M': 1 << 20, // Мегабайт
		'G': 1 << 30, // Гигабайт
		'T': 1 << 40, // Терабайт
		'P': 1 << 50, // Петабайт
	}

	// Если строка пустая, вернуть ошибку
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string")
	}

	// Проверить, есть ли суффикс
	lastChar := s[len(s)-1]
	if multiplier, ok := units[lastChar]; ok {
		value, err := strconv.ParseFloat(s[:len(s)-1], 64)
		if err != nil {
			return 0, err
		}
		return value * multiplier, nil
	}

	// Если нет суффикса, попытаться преобразовать строку в число
	return strconv.ParseFloat(s, 64)
}

// Основная функция сортировки файла
func sortFile() {
	// Определение флагов командной строки
	k := flag.Int("k", 0, "указание колонки для сортировки")
	n := flag.Bool("n", false, "сортировать по числовому значению")
	r := flag.Bool("r", false, "сортировать в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	M := flag.Bool("M", false, "сортировать по названию месяца")
	b := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	c := flag.Bool("c", false, "проверять отсортированы ли данные")
	h := flag.Bool("h", false, "сортировать по числовому значению с учётом суффиксов")

	flag.Parse()

	// Проверка наличия входного файла
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: sort [options] inputfile")
		return
	}
	inputFile := args[0]

	// Открытие входного файла
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Чтение строк из файла
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if *b {
			line = strings.TrimSpace(line) // Удаление хвостовых пробелов, если указан флаг -b
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Функция для получения ключа сортировки из строки
	getKey := func(line string) string {
		if *k > 0 {
			fields := strings.Fields(line)
			if *k-1 < len(fields) {
				return fields[*k-1]
			}
		}
		return line
	}

	// Сортировка строк
	sort.Slice(lines, func(i, j int) bool {
		a, b := getKey(lines[i]), getKey(lines[j])
		if *n {
			ai, err1 := strconv.ParseFloat(a, 64)
			bi, err2 := strconv.ParseFloat(b, 64)
			if err1 == nil && err2 == nil {
				if *r {
					return ai > bi
				}
				return ai < bi
			}
		}
		if *h {
			ai, err1 := parseHumanReadableNumber(a)
			bi, err2 := parseHumanReadableNumber(b)
			if err1 == nil && err2 == nil {
				if *r {
					return ai > bi
				}
				return ai < bi
			}
		}
		if *M {
			months := map[string]int{
				"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6,
				"July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12,
			}
			ai, ok1 := months[a]
			bi, ok2 := months[b]
			if ok1 && ok2 {
				if *r {
					return ai > bi
				}
				return ai < bi
			}
		}
		if *r {
			return a > b
		}
		return a < b
	})

	// Удаление повторяющихся строк, если указан флаг -u
	if *u {
		uniqueLines := []string{}
		seen := map[string]bool{}
		for _, line := range lines {
			if !seen[line] {
				uniqueLines = append(uniqueLines, line)
				seen[line] = true
			}
		}
		lines = uniqueLines
	}

	// Проверка отсортированы ли данные, если указан флаг -c
	if *c {
		sorted := sort.SliceIsSorted(lines, func(i, j int) bool {
			a, b := getKey(lines[i]), getKey(lines[j])
			if *n {
				ai, err1 := strconv.ParseFloat(a, 64)
				bi, err2 := strconv.ParseFloat(b, 64)
				if err1 == nil && err2 == nil {
					return ai < bi
				}
			}
			if *h {
				ai, err1 := parseHumanReadableNumber(a)
				bi, err2 := parseHumanReadableNumber(b)
				if err1 == nil && err2 == nil {
					return ai < bi
				}
			}
			if *M {
				months := map[string]int{
					"January": 1, "February": 2, "March": 3, "April": 4, "May": 5, "June": 6,
					"July": 7, "August": 8, "September": 9, "October": 10, "November": 11, "December": 12,
				}
				ai, ok1 := months[a]
				bi, ok2 := months[b]
				if ok1 && ok2 {
					return ai < bi
				}
			}
			return a < b
		})
		if sorted {
			fmt.Println("Файл отсортирован")
		} else {
			fmt.Println("Файл не отсортирован")
		}
		return
	}

	// Запись отсортированных строк в файл
	outFile, err := os.Create(inputFile)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
}

func main() {
	sortFile()
}
