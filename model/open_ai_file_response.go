package model

type OpenAIFileResponse struct {
	Data   []OpenAIFile `json:"data"`
	Object string       `json:"object"`
}