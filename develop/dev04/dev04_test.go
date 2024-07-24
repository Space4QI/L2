package main

import (
	"reflect"
	"testing"
)

// Тестовая функция для поиска множеств анаграмм
func TestAnagram(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected map[string][]string
	}{
		{
			name:  "Basic anagrams",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "Single word",
			words:    []string{"пятак"},
			expected: map[string][]string{},
		},
		{
			name:  "Case insensitivity",
			words: []string{"кино", "икон", "кИно"},
			expected: map[string][]string{
				"икон": {"икон", "кино", "кино"},
			},
		},
		{
			name:     "No anagrams",
			words:    []string{"мир", "земля", "небо"},
			expected: map[string][]string{},
		},
		{
			name:  "Multiple groups of anagrams",
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := anagram(tt.words)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
