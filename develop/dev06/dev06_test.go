package main

import (
	"testing"
)

func TestCut(t *testing.T) {
	tests := []struct {
		name     string
		options  cutOptions
		input    []string
		expected []string
	}{
		{
			name: "Basic cut",
			options: cutOptions{
				fields:    "1",
				delimiter: "\t",
			},
			input:    []string{"col1\tcol2\tcol3", "val1\tval2\tval3"},
			expected: []string{"col1", "val1"},
		},
		{
			name: "Multiple fields",
			options: cutOptions{
				fields:    "1,3",
				delimiter: "\t",
			},
			input:    []string{"col1\tcol2\tcol3", "val1\tval2\tval3"},
			expected: []string{"col1\tcol3", "val1\tval3"},
		},
		{
			name: "Different delimiter",
			options: cutOptions{
				fields:    "2",
				delimiter: ",",
			},
			input:    []string{"col1,col2,col3", "val1,val2,val3"},
			expected: []string{"col2", "val2"},
		},
		{
			name: "Separated lines only",
			options: cutOptions{
				fields:    "2",
				delimiter: "\t",
				separated: true,
			},
			input:    []string{"col1\tcol2\tcol3", "no tabs here", "val1\tval2\tval3"},
			expected: []string{"col2", "val2"},
		},
		{
			name: "No fields specified",
			options: cutOptions{
				fields:    "",
				delimiter: "\t",
			},
			input:    []string{"col1\tcol2\tcol3", "val1\tval2\tval3"},
			expected: []string{"col1\tcol2\tcol3", "val1\tval2\tval3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cut(tt.options, tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d results, got %d", len(tt.expected), len(result))
			}
			for i, line := range result {
				if line != tt.expected[i] {
					t.Errorf("expected '%s', got '%s'", tt.expected[i], line)
				}
			}
		})
	}
}
