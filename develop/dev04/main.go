package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagram(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	// Пройти по всем словам
	for _, word := range words {
		// Привести слово к нижнему регистру
		lowerWord := strings.ToLower(word)

		// Получить ключ анаграммы: отсортировать буквы в слове
		sortedWord := sortStringByCharacter(lowerWord)

		// Добавить слово в соответствующее множество анаграмм
		anagramMap[sortedWord] = append(anagramMap[sortedWord], lowerWord)
	}

	// Создаём мапу для результата
	result := make(map[string][]string)

	// Проходим по мапе анаграмм и заполняем
	for _, anagrams := range anagramMap {
		if len(anagrams) > 1 {
			sort.Strings(anagrams)
			result[anagrams[0]] = anagrams
		}
	}

	return result
}

// Доп. функция для сортировки букв в строке
func sortStringByCharacter(str string) string {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}

	anagrams := anagram(words)

	for key, group := range anagrams {
		fmt.Printf("%s: %v\n", key, group)
	}
}
