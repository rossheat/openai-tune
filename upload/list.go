package upload

import (
    "encoding/json"
    "fmt"
    "net/url"
    "strings"
    "time"

    "github.com/rossheat/openai-tune/http"
    "github.com/rossheat/openai-tune/model"
    "github.com/rossheat/openai-tune/option"
)

func List(options option.Upload) error {
    params := url.Values{}
    params.Add("purpose", "fine-tune")
    params.Add("order", "desc")
    params.Add("limit", "10000")

    client := http.NewClient(options.OpenAIAPIKey)
    body, err := client.DoWithParams("GET", "/files", params, nil)
    if err != nil {
        return err
    }

    var fileResp model.OpenAIFileResponse
    if err := json.Unmarshal(body, &fileResp); err != nil {
        return fmt.Errorf("error parsing response: %v", err)
    }

    fmt.Println("Files uploaded for fine-tuning:")
    fmt.Printf("%-30s %-20s %-15s %-20s\n", "ID", "Filename", "Size", "Uploaded")
    fmt.Printf("%-30s %-20s %-15s %-20s\n", 
        strings.Repeat("-", 30), 
        strings.Repeat("-", 20), 
        strings.Repeat("-", 15), 
        strings.Repeat("-", 20))

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