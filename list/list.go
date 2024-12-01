package list

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rossheat/openai-tune/option"
)

func List(options option.List) error {
	baseURL := "https://api.openai.com/v1/fine_tuning/jobs"

	params := url.Values{}
	if options.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", options.Limit))
	}
	if options.After != "" {
		params.Add("after", options.After)
	}

	requestURL := baseURL
	if len(params) > 0 {
		requestURL += "?" + params.Encode()
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+options.OpenAIAPIKey)

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

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("error parsing response: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		return fmt.Errorf("error formatting JSON: %v", err)
	}

	fmt.Println(string(prettyJSON))

	if hasMore, ok := response["has_more"].(bool); ok && hasMore {
		fmt.Println("\nThere are more results available. Use the -after flag with the last job ID to see more.")
	}

	return nil
}
