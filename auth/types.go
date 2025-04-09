package main

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
	AccessToken string `json:"access_token"`
}

type KeycloakService interface {
	login(payload *KLoginPayload) error
}
