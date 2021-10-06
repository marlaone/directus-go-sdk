package directusgosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AuthInfo struct {
	AccessToken  string `json:"access_token"`
	Expires      string `json:"expires"`
	RefreshToken string `json:"refresh_token"`
}

type Client struct {
	*Config
	AuthInfo
}

func NewClient(c *Config) *Client {
	return &Client{
		Config: c,
		AuthInfo: AuthInfo{
			AccessToken:  "",
			Expires:      "",
			RefreshToken: "",
		},
	}
}

func (c *Client) Login(email, password string) error {
	loginBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return fmt.Errorf("parsing login body error: %v", err)
	}

	authResponse, err := NewDirectusRequest(c, "/auth/login", "POST", bytes.NewBuffer(loginBody))

	if err != nil {
		return fmt.Errorf("login request error: %v", err)
	}

	authResponseData, ok := authResponse.Data.(map[string]interface{})

	if !ok && authResponse.Errors == nil {
		return fmt.Errorf("response data invalid type")
	}

	if accessToken, ok := authResponseData["access_token"].(string); ok {
		c.AuthInfo = AuthInfo{
			AccessToken:  accessToken,
			Expires:      fmt.Sprintf("%f", authResponseData["expires"].(float64)),
			RefreshToken: authResponseData["refresh_token"].(string),
		}
	} else if authResponse.Errors != nil && len(authResponse.Errors) > 0 {
		return fmt.Errorf("login error: %v", authResponse.Errors)
	} else {
		return fmt.Errorf("invalid auth response type")
	}

	return nil
}
