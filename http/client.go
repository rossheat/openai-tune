package http

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
)

type Client struct {
    apiKey     string
    baseURL    string
    httpClient *http.Client
}

func NewClient(apiKey string) *Client {
    return &Client{
        apiKey:     apiKey,
        baseURL:    "https://api.openai.com/v1",
        httpClient: &http.Client{},
    }
}

func (c *Client) Do(method, path string, body interface{}) ([]byte, error) {
    var reqBody io.Reader
    if body != nil {
        jsonData, err := json.Marshal(body)
        if err != nil {
            return nil, fmt.Errorf("error marshaling request body: %v", err)
        }
        reqBody = bytes.NewBuffer(jsonData)
    }

    url := c.baseURL + path
    req, err := http.NewRequest(method, url, reqBody)
    if err != nil {
        return nil, fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Authorization", "Bearer "+c.apiKey)
    if body != nil {
        req.Header.Set("Content-Type", "application/json")
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
    }

    return respBody, nil
}

func (c *Client) DoWithParams(method, path string, params url.Values, body interface{}) ([]byte, error) {
    if len(params) > 0 {
        path += "?" + params.Encode()
    }
    return c.Do(method, path, body)
}

func (c *Client) DoMultipart(path, contentType string, body io.Reader) ([]byte, error) {
    url := c.baseURL + path
    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return nil, fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Authorization", "Bearer "+c.apiKey)
    req.Header.Set("Content-Type", contentType)

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
    }

    return respBody, nil
}