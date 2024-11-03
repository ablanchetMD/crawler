package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestPrintReports(t *testing.T) {
	tests := []struct {
		name     string
		pages    map[string]int
		baseURL  string
		expected string
	}{
		{
			name:    "first test",
			pages:   map[string]int{"https://blog.boot.dev/path/one": 1, "https://other.com/path/one": 1},
			baseURL: "https://blog.boot.dev",
			expected: "=============================\n" +
				"  REPORT for https://blog.boot.dev\n" +
				"=============================\n" +
				"Found 1 internal links to https://blog.boot.dev/path/one\n" +
				"Found 1 internal links to https://other.com/path/one\n",
		},
		{
			name:    "2nd test",
			pages:   map[string]int{"https://blog.boot.dev/path/one": 3, "https://other.com/path/one": 1, "https://blog.boot.dev/path/two": 2, "https://blog.boot.dev/path/three": 1},
			baseURL: "https://blog.boot.dev",
			expected: "=============================\n" +
				"  REPORT for https://blog.boot.dev\n" +
				"=============================\n" +
				"Found 3 internal links to https://blog.boot.dev/path/one\n" +
				"Found 2 internal links to https://blog.boot.dev/path/two\n" +
				"Found 1 internal links to https://blog.boot.dev/path/three\n" +
				"Found 1 internal links to https://other.com/path/one\n",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			originalStdout := os.Stdout
			os.Stdout = w

			var buf bytes.Buffer

			PrintReports(tc.pages, tc.baseURL)

			w.Close()
			os.Stdout = originalStdout

			io.Copy(&buf, r)
			r.Close()

			if !reflect.DeepEqual(buf.String(), tc.expected) {
				t.Errorf("Test %v - %s FAIL", i, tc.name)
				fmt.Println("Expected:")
				fmt.Print(tc.expected)
				fmt.Println("Actual:")
				fmt.Print(buf.String())
			}
		})
	}
}
