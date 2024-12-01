package upload

import (
	"fmt"

	"bufio"
	"encoding/json"
	"os"

	"github.com/rossheat/openai-tune/options"
)

func isValidJSONL(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var jsonObj interface{}
		if err := json.Unmarshal([]byte(line), &jsonObj); err != nil {
			return lineCount, fmt.Errorf("invalid JSONL file: %v", err)
		}
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return lineCount, err
	}

	if lineCount == 0 {
		return 0, fmt.Errorf("invalid JSONL file: no JSON lines found")
	}

	return lineCount, nil
}

func Upload(options options.Upload) error {
	fmt.Printf("Upload function received options: %v\n", options)

	lines, err := isValidJSONL(options.File)
	if err != nil {
		return err
	}
	fmt.Printf("Valid JSONL file with %d lines\n", lines)

	return nil
}
