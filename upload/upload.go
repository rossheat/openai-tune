package upload

import (
	"fmt"

	"bufio"
	"encoding/json"
	"os"

	"github.com/rossheat/openai-tune/options"
)

func isValidJSONL(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		var jsonObj interface{}
		if err := json.Unmarshal([]byte(line), &jsonObj); err != nil {
			return fmt.Errorf("invalid JSONL file: %v", err)
		}
	}

	return scanner.Err()
}

func Upload(options options.Upload) error {
	fmt.Printf("Upload function received options: %v\n", options)

	err := isValidJSONL(options.File)
	if err != nil {
		return err
	}

	return nil
}
