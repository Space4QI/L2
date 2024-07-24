package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSortFile(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		inputContent string
		expected     string
	}{
		{
			name:         "NumericSort",
			args:         []string{"-n"},
			inputContent: "10\n5\n20\n2\n",
			expected:     "2\n5\n10\n20\n",
		},
		{
			name:         "ReverseSort",
			args:         []string{"-r"},
			inputContent: "apple\norange\nbanana\ngrape\n",
			expected:     "orange\ngrape\nbanana\napple\n",
		},
		{
			name:         "UniqueSort",
			args:         []string{"-u"},
			inputContent: "apple\norange\nbanana\napple\n",
			expected:     "apple\nbanana\norange\n",
		},
		{
			name:         "MonthSort",
			args:         []string{"-M"},
			inputContent: "July\nJanuary\nMarch\nFebruary\n",
			expected:     "January\nFebruary\nMarch\nJuly\n",
		},
		{
			name:         "HumanReadableSort",
			args:         []string{"-h"},
			inputContent: "100K\n50K\n1G\n500M\n",
			expected:     "50K\n100K\n500M\n1G\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpfile, err := ioutil.TempFile("", "testfile")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name()) // clean up

			if _, err := tmpfile.WriteString(tt.inputContent); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			os.Args = append([]string{"cmd"}, tt.args...)
			os.Args = append(os.Args, tmpfile.Name())

			sortFile()

			sortedContent, err := ioutil.ReadFile(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}

			if string(sortedContent) != tt.expected {
				t.Errorf("expected:\n%s\ngot:\n%s\n", tt.expected, string(sortedContent))
			}
		})
	}
}
