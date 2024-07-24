package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpacking(input string) string {
	// Преобразуем входную строку в массив рун
	str := []rune(input)

	// Создаем буфер для сохранения распакованной строки
	var result strings.Builder

	// Проходим по каждому символу в строке
	for i := 0; i < len(str); i++ {
		// Проверяем, если первый символ является цифрой
		if unicode.IsDigit(str[0]) {
			fmt.Println("Ошибка: некорректная строка")
			return ""
		}

		// Проверяем, если предпоследний символ не экранирован, а последний экранирован
		if str[len(str)-2] != '\\' && str[len(str)-1] == '\\' {
			fmt.Println("Ошибка: некорректная строка")
			return ""
		}

		// Если текущий символ экранированный
		if str[i] == '\\' {
			// Добавляем следующий символ к результату
			result.WriteString(string(str[i+1]))

			// Инкрементируем i, чтобы пропустить следующий символ
			i++
		} else {
			// Если текущий символ не экранированный

			// Если текущий символ является цифрой
			if unicode.IsDigit(str[i]) {
				// Конвертируем его в число
				count, _ := strconv.Atoi(string(str[i]))

				// Добавляем предыдущий символ count-1 раз к результату
				result.WriteString(strings.Repeat(string(str[i-1]), count-1))
			} else {
				// Если текущий символ не является цифрой и не экранированным символом,
				// просто добавляем его к результату
				result.WriteString(string(str[i]))
			}
		}
	}

	// Возвращаем распакованную строку
	return result.String()
}

func main() {
	// Входные данные
	s := ""

	// Распаковка строки
	result := unpacking(s)
	// Вывод результата
	fmt.Println(result)

}
