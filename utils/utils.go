package utils

import (
	"fmt"
	"os"
)

func GetOpenAIAPIKeyFromEnv() (string, error) {
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if openAIAPIKey == "" {
		return "", fmt.Errorf("failed to get OPENAI_API_KEY from environment; set it like so: export=OPENAI_API_KEY=<OPENAI_API_KEY>")
	}
	return openAIAPIKey, nil
}
