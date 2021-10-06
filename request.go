package directusgosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DirectusErrors struct {
	Message    string
	Extensions map[string]string
}

type DirectusResponse struct {
	Data   interface{}
	Errors []DirectusErrors
}

func BuildRequestUri(c *Config, url string) string {
	return fmt.Sprintf("%s%s", c.GetEndpoint(), url)
}

func NewDirectusRequest(c *Client, url string, method string, requestBody io.Reader) (*DirectusResponse, error) {
	var directusResponse *DirectusResponse
	client := &http.Client{}

	req, err := http.NewRequest(method, BuildRequestUri(c.Config, url), requestBody)
	if err != nil {
		return nil, fmt.Errorf("creating request error: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	if c.AccessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("directus request error: %v", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&directusResponse); err != nil {
		return nil, fmt.Errorf("directus response error: %v", err)
	}

	return directusResponse, nil
}
