package directusgosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AuthResponse struct {
	Data   *AuthInfo         `json:"data"`
	Errors []*DirectusErrors `json:"errors"`
}

type AuthInfo struct {
	AccessToken  string  `json:"access_token"`
	Expires      float64 `json:"expires"`
	RefreshToken string  `json:"refresh_token"`
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
			Expires:      0,
			RefreshToken: "",
		},
	}
}

func (c *Client) Login(email, password string) error {
	var authResponse AuthResponse

	loginBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return fmt.Errorf("parsing login body error: %v", err)
	}

	resp, err := NewDirectusRequest(c, "/auth/login", "POST", bytes.NewBuffer(loginBody))

	if err != nil {
		return fmt.Errorf("login request error: %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return fmt.Errorf("login response error: %v", err)
	}

	if authResponse.Data != nil {
		c.AuthInfo = *authResponse.Data
	} else if authResponse.Errors != nil && len(authResponse.Errors) > 0 {
		return fmt.Errorf("login error: %v", authResponse.Errors)
	} else {
		return fmt.Errorf("invalid auth response type")
	}

	return nil
}
