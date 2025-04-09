package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func NewKCClient() *KCClient {
	return &KCClient{
		httpClient: http.DefaultClient,
	}
}
func (c *KCClient) login(payload *KLoginPayload) (*KLoginRes, error) {
	formData := url.Values{
		"client_id":     {payload.clientId},
		"username":      {payload.userName},
		"password":      {payload.password},
		"grant_type":    {payload.grantType},
		"client_secret": {payload.clientSecret},
	}
	encodedFormData := formData.Encode()
	req, err := http.NewRequest("POST", "http://localhost:8080/realms/rest-golang/protocol/openid-connect/token", strings.NewReader(encodedFormData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed with status code: %d: %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	result := &KLoginRes{}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return result, nil
}
