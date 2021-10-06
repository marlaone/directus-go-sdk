package directusgosdk

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DirectusErrors struct {
	Message    string
	Extensions map[string]string
}

type DirectusResponse struct {
	Data   interface{}
	Errors []DirectusErrors
}

func BuildRequestUri(c *Config, url string, q *Query) string {
	uri := fmt.Sprintf("%s%s", c.GetEndpoint(), url)

	if q != nil {
		uri = uri + "?" + q.ToQueryString()
	}

	return uri
}

func NewDirectusRequest(c *Client, url string, method string, requestBody io.Reader, q *Query) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, BuildRequestUri(c.Config, url, q), requestBody)
	if err != nil {
		return nil, fmt.Errorf("creating request error: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")

	if c.AccessToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	}

	return client.Do(req)
}

func CheckDirectusErrors(errors []DirectusErrors) error {
	if len(errors) > 0 {
		errorMessages := []string{}
		for _, err := range errors {
			errorMessages = append(errorMessages, err.Message)
		}

		return fmt.Errorf(strings.Join(errorMessages, ", "))
	}

	return nil
}
