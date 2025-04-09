package main

import "net/http"

// GreetRes is the response structure for the greeting endpoint.
type GreetRes struct {
	Hello string `json:"hello"`
}

type KLoginPayload struct {
	clientId     string
	userName     string
	password     string
	grantType    string
	clientSecret string
}

type KLoginRes struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type KeycloakService interface {
	login(payload *KLoginPayload) (*KLoginRes, error)
}

type KCClient struct {
	httpClient *http.Client
}
