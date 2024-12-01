package model

type OpenAIFile struct {
	ID             string `json:"id"`
	Object         string `json:"object"`
	Bytes          int    `json:"bytes"`
	CreatedAt      int64  `json:"created_at"`
	OpenAIFilename string `json:"filename"`
	Purpose        string `json:"purpose"`
}

