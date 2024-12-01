package cancel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rossheat/openai-tune/option"
)

func Cancel(options option.Job) error {
	url := fmt.Sprintf("https://api.openai.com/v1/fine_tuning/jobs/%s/cancel", options.JobID)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+options.OpenAIAPIKey)

	client := &http.Client{}
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

	var prettyJSON map[string]interface{}
	if err := json.Unmarshal(body, &prettyJSON); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	prettyOutput, err := json.MarshalIndent(prettyJSON, "", "    ")
	if err != nil {
		return fmt.Errorf("error formatting JSON: %v", err)
	}

	fmt.Println(string(prettyOutput))
	return nil
}
