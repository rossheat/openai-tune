package upload

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/files", &buf)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+options.OpenAIAPIKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}
	fmt.Printf("Upload successful: %s\n", string(body))

	return nil
}
