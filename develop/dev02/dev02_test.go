package main

import (
	"strconv"
	"testing"
)

func TestUnpacking(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		hasError bool
	}{
		// Успешные случаи
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},

		// Случаи с ошибкой
		{"45", "", true},
		{"qwe\\\\5\\", "", true},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			result := unpacking(tc.input)

			if result != tc.expected {
				t.Errorf("Test case %d: ожидалось \"%s\", получено %s", i+1, tc.expected, result)
			}

			// Проверка ошибки, если ожидается ошибка и она не возникла, или наоборот
			if tc.hasError && result != "" || !tc.hasError && result == "" {
				t.Errorf("Test case %d: ожидалась ошибка, результат %v", i+1, result)
			}
		})
	}
}
