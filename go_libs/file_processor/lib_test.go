package fileprocessor

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Helper function to create a temporary test file
func createTestFile(t *testing.T, content string) string {
	t.Helper()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test_input.txt")

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	return filePath
}

func TestLoadRaw(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{
			name:     "simple text",
			content:  "hello world",
			expected: "hello world",
		},
		{
			name:     "multiline text",
			content:  "line1\nline2\nline3",
			expected: "line1\nline2\nline3",
		},
		{
			name:     "empty file",
			content:  "",
			expected: "",
		},
		{
			name:     "text with special characters",
			content:  "hello\tworld\r\n!@#$%",
			expected: "hello\tworld\r\n!@#$%",
		},
	}

	fp := &fileProcessorImpl{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := createTestFile(t, tt.content)
			result := fp.LoadRaw(filePath)

			if result != tt.expected {
				t.Errorf("LoadRaw() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestLoadLines(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []string
	}{
		{
			name:     "single line",
			content:  "hello",
			expected: []string{"hello"},
		},
		{
			name:     "multiple lines",
			content:  "line1\nline2\nline3",
			expected: []string{"line1", "line2", "line3"},
		},
		{
			name:     "empty file",
			content:  "",
			expected: []string{""},
		},
		{
			name:     "lines with empty lines",
			content:  "line1\n\nline3",
			expected: []string{"line1", "", "line3"},
		},
		{
			name:     "trailing newline",
			content:  "line1\nline2\n",
			expected: []string{"line1", "line2", ""},
		},
	}

	fp := &fileProcessorImpl{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := createTestFile(t, tt.content)
			result := fp.LoadLines(filePath)

			if len(result) != len(tt.expected) {
				t.Errorf("LoadLines() returned %d lines, want %d", len(result), len(tt.expected))
				return
			}

			for i, line := range result {
				if line != tt.expected[i] {
					t.Errorf("LoadLines()[%d] = %q, want %q", i, line, tt.expected[i])
				}
			}
		})
	}
}

func TestProcessLineByLine(t *testing.T) {
	tests := []struct {
		name          string
		content       string
		processFunc   func(string) string
		expectedCalls []string
	}{
		{
			name:    "process multiple lines",
			content: "line1\nline2\nline3",
			processFunc: func(s string) string {
				return strings.ToUpper(s)
			},
			expectedCalls: []string{"line1\n", "line2\n", "line3"},
		},
		{
			name:    "process single line",
			content: "single",
			processFunc: func(s string) string {
				return s
			},
			expectedCalls: []string{"single"},
		},
		{
			name:    "process with transformation",
			content: "a\nb\nc",
			processFunc: func(s string) string {
				return strings.TrimSpace(s) + "!"
			},
			expectedCalls: []string{"a\n", "b\n", "c"},
		},
	}

	fp := &fileProcessorImpl{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := createTestFile(t, tt.content)

			var calls []string
			captureFunc := func(line string) string {
				calls = append(calls, line)
				return tt.processFunc(line)
			}

			fp.ProcessLineByLine(filePath, captureFunc)

			if len(calls) != len(tt.expectedCalls) {
				t.Errorf("ProcessLineByLine() called processFunc %d times, want %d", len(calls), len(tt.expectedCalls))
				return
			}

			for i, call := range calls {
				if call != tt.expectedCalls[i] {
					t.Errorf("ProcessLineByLine() call[%d] = %q, want %q", i, call, tt.expectedCalls[i])
				}
			}
		})
	}
}

func TestProcessLineByLine_EmptyFile(t *testing.T) {
	fp := &fileProcessorImpl{}
	filePath := createTestFile(t, "")

	callCount := 0
	processFunc := func(line string) string {
		callCount++
		return line
	}

	fp.ProcessLineByLine(filePath, processFunc)

	if callCount != 0 {
		t.Errorf("ProcessLineByLine() called processFunc %d times for empty file, want 0", callCount)
	}
}

// Benchmark tests
func BenchmarkLoadRaw(b *testing.B) {
	content := strings.Repeat("test line\n", 1000)
	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "bench_test.txt")
	os.WriteFile(filePath, []byte(content), 0644)

	fp := &fileProcessorImpl{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fp.LoadRaw(filePath)
	}
}

func BenchmarkLoadLines(b *testing.B) {
	content := strings.Repeat("test line\n", 1000)
	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "bench_test.txt")
	os.WriteFile(filePath, []byte(content), 0644)

	fp := &fileProcessorImpl{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fp.LoadLines(filePath)
	}
}

func BenchmarkProcessLineByLine(b *testing.B) {
	content := strings.Repeat("test line\n", 1000)
	tmpDir := b.TempDir()
	filePath := filepath.Join(tmpDir, "bench_test.txt")
	os.WriteFile(filePath, []byte(content), 0644)

	fp := &fileProcessorImpl{}
	processFunc := func(line string) string {
		return strings.ToUpper(line)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fp.ProcessLineByLine(filePath, processFunc)
	}
}
