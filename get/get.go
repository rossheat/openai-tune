package get

import (
    "encoding/json"
    "fmt"

    "github.com/rossheat/openai-tune/http"
    "github.com/rossheat/openai-tune/option"
)

func Get(options option.Job) error {
    client := http.NewClient(options.OpenAIAPIKey)
    body, err := client.Do("GET", fmt.Sprintf("/fine_tuning/jobs/%s", options.JobID), nil)
    if err != nil {
        return err
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