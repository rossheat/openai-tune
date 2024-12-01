// create/create.go
package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rossheat/openai-tune/model"
	"github.com/rossheat/openai-tune/option"
	"gopkg.in/yaml.v3"
)

func Create(options option.Create) error {
	var config model.FineTuneConfig

	if options.ConfigFile != "" {
		data, err := os.ReadFile(options.ConfigFile)
		if err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}

		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return fmt.Errorf("error parsing config file: %v", err)
		}
	} else {
		config = model.FineTuneConfig{
			Model:        options.Model,
			TrainingFile: options.FileID,
		}
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/fine_tuning/jobs",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+options.OpenAIAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Println(string(body))
	return nil
}
