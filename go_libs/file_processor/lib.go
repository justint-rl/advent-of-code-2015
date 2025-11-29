package fileprocessor

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type FileProcessor interface {
	LoadRaw(filePath string) string
	LoadLines(filePath string) []string
	ProcessLineByLine(filePath string, processFunc func(string) string)
}

type fileProcessorImpl struct{}

func New() FileProcessor {
	return &fileProcessorImpl{}
}

func (f *fileProcessorImpl) LoadRaw(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
		return ""
	}

	return string(data)
}

func (f *fileProcessorImpl) LoadLines(filePath string) []string {
	data := f.LoadRaw(filePath)

	lines := strings.Split(data, "\n")
	return lines
}

func (f *fileProcessorImpl) ProcessLineByLine(filePath string, processFunc func(string) string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // reach to new line
		if err != nil {
			if err == io.EOF {
				// Process the last line if there's any data
				if len(line) > 0 {
					processFunc(line)
				}
				break // Reached end of file
			}
			log.Fatalf("Error reading line in file: %s", err)
		}
		processFunc(line)
	}
}
