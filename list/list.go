package list

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/rossheat/openai-tune/http"
	"github.com/rossheat/openai-tune/option"
)

func List(options option.List) error {
	params := url.Values{}
	if options.Limit > 0 {
		params.Add("limit", fmt.Sprintf("%d", options.Limit))
	}
	if options.After != "" {
		params.Add("after", options.After)
	}

	client := http.NewClient(options.OpenAIAPIKey)
	body, err := client.DoWithParams("GET", "/fine_tuning/jobs", params, nil)
	if err != nil {
		return err
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
