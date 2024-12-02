package upload

import (
    "encoding/json"
    "fmt"
    "net/url"
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

    prettyJSON, err := json.MarshalIndent(fileResp, "", "    ")
    if err != nil {
        return fmt.Errorf("error formatting JSON: %v", err)
    }

    fmt.Println(string(prettyJSON))
    return nil
}