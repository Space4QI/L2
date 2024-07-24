package main

import (
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name     string
		options  grepOptions
		input    []string
		expected []string
	}{
		{
			name: "Basic match",
			options: grepOptions{
				pattern: "test",
			},
			input:    []string{"this is a test", "another line"},
			expected: []string{"this is a test"},
		},
		{
			name: "Ignore case",
			options: grepOptions{
				pattern:    "Test",
				ignoreCase: true,
			},
			input:    []string{"this is a test", "another line"},
			expected: []string{"this is a test"},
		},
		{
			name: "Invert match",
			options: grepOptions{
				pattern: "test",
				invert:  true,
			},
			input:    []string{"this is a test", "another line"},
			expected: []string{"another line"},
		},
		{
			name: "Fixed string match",
			options: grepOptions{
				pattern: "test",
				fixed:   true,
			},
			input:    []string{"test", "testing", "another test"},
			expected: []string{"test"},
		},
		{
			name: "Count matches",
			options: grepOptions{
				pattern: "test",
				count:   true,
			},
			input:    []string{"this is a test", "another test"},
			expected: []string{"2"},
		},
		{
			name: "Line numbers",
			options: grepOptions{
				pattern: "test",
				lineNum: true,
			},
			input:    []string{"this is a test", "another line", "test line"},
			expected: []string{"1:this is a test", "3:test line"},
		},
		{
			name: "Before context",
			options: grepOptions{
				pattern: "test",
				before:  1,
			},
			input:    []string{"line before", "this is a test", "line after"},
			expected: []string{"line before", "this is a test"},
		},
		{
			name: "After context",
			options: grepOptions{
				pattern: "test",
				after:   1,
			},
			input:    []string{"line before", "this is a test", "line after"},
			expected: []string{"this is a test", "line after"},
		},
		{
			name: "Context",
			options: grepOptions{
				pattern: "test",
				before:  1,
				after:   1,
			},
			input:    []string{"line before", "this is a test", "line after"},
			expected: []string{"line before", "this is a test", "line after"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := grep(tt.options, tt.input)
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
