package upload

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/rossheat/openai-tune/http"
	"github.com/rossheat/openai-tune/option"
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
		lineCount++

		var data struct {
			Messages []struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"messages"`
		}

		if err := json.Unmarshal([]byte(line), &data); err != nil {
			return lineCount, fmt.Errorf("line %d: invalid JSON: %v", lineCount, err)
		}

		if data.Messages[0].Role != "system" {
			return lineCount, fmt.Errorf("line %d: first message must have role 'system'", lineCount)
		}

		if len(data.Messages) > 1 {
			if data.Messages[1].Role != "user" {
				return lineCount, fmt.Errorf("line %d: second message must have role 'user'", lineCount)
			}

			for i, msg := range data.Messages[2:] {
				if msg.Role != "assistant" && msg.Role != "user" {
					return lineCount, fmt.Errorf("line %d: message %d must have role 'assistant' or 'user'",
						lineCount, i+3)
				}
				if msg.Role == "user" && data.Messages[i+1].Role == "user" {
					return lineCount, fmt.Errorf("line %d: cannot have consecutive user messages", lineCount)
				}
				if msg.Role == "assistant" && data.Messages[i+1].Role == "assistant" {
					return lineCount, fmt.Errorf("line %d: cannot have consecutive assistant messages", lineCount)
				}
			}
		}

		for i, msg := range data.Messages {
			if msg.Content == "" {
				return lineCount, fmt.Errorf("line %d: message %d has empty content", lineCount, i+1)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return lineCount, err
	}

	if lineCount == 0 {
		return 0, fmt.Errorf("invalid JSONL file: no JSON lines found")
	}

	return lineCount, nil
}

func Upload(options option.Upload) error {
	fmt.Printf("Upload function received options: %v\n", options)

	lines, err := isValidJSONL(options.File)
	if err != nil {
		return err
	}
	fmt.Printf("Valid JSONL file with %d lines\n", lines)

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	if err := writer.WriteField("purpose", "fine-tune"); err != nil {
		return fmt.Errorf("error writing purpose field: %v", err)
	}

	file, err := os.Open(options.File)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(options.File))
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}

	if _, err = io.Copy(part, file); err != nil {
		return fmt.Errorf("error copying file content: %v", err)
	}

	writer.Close()

	client := http.NewClient(options.OpenAIAPIKey)
	body, err := client.DoMultipart("/files", writer.FormDataContentType(), &buf)
	if err != nil {
		return err
	}

	fmt.Printf("Upload successful: %s\n", string(body))
	return nil
}
