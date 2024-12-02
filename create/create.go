package create

import (
	"fmt"
	"os"

	"github.com/rossheat/openai-tune/http"
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

	client := http.NewClient(options.OpenAIAPIKey)
	body, err := client.Do("POST", "/fine_tuning/jobs", config)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
