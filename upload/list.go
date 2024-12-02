package upload

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rossheat/openai-tune/model"
	"github.com/rossheat/openai-tune/option"
)

func List(options option.Upload) error {
	baseURL, err := url.Parse("https://api.openai.com/v1/files")
	if err != nil {
		return fmt.Errorf("error parsing URL: %v", err)
	}

	params := url.Values{}
	params.Add("purpose", "fine-tune")
	params.Add("order", "desc")
	params.Add("limit", "10000")
	baseURL.RawQuery = params.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+options.OpenAIAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var fileResp model.OpenAIFileResponse
	if err := json.Unmarshal(body, &fileResp); err != nil {
		return fmt.Errorf("error parsing response: %v", err)
	}

	fmt.Println("Files uploaded for fine-tuning:")
	fmt.Printf("%-30s %-20s %-15s %-20s\n", "ID", "Filename", "Size", "Uploaded")
	fmt.Printf("%-30s %-20s %-15s %-20s\n", strings.Repeat("-", 30), strings.Repeat("-", 20), strings.Repeat("-", 15), strings.Repeat("-", 20))

	for _, file := range fileResp.Data {
		uploadTime := time.Unix(file.CreatedAt, 0).Format("02/01/2006 15:04:05")
		fmt.Printf("%-30s %-20s %-15s %-20s\n",
			file.ID,
			file.OpenAIFilename,
			fmt.Sprintf("%d bytes", file.Bytes),
			uploadTime)
	}

	return nil
}
