package entity

import (
	"time"
)

type RequestToken struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
}

type Token struct {
	TokenType    string    `json:"token_type"`
	ExpireIn     time.Time `json:"expires_in"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
