package internal

import (
	"context"

	oauth2 "golang.org/x/oauth2"
	cc "golang.org/x/oauth2/clientcredentials"
)

const ExampleTokenHmacSecret = "my_secret"

type (
	OAuth2Config struct {
		ClientID       string
		ClientSecret   string
		TokenEndepoint string
	}
)

var (
	Port    int
	Channel string
	OAuth2  OAuth2Config

	tokenSource oauth2.TokenSource
)

type ChatMessage struct {
	Input string `json:"input"`
}

func GetTokenSource(ctx context.Context) oauth2.TokenSource {
	if tokenSource == nil {
		config := &cc.Config{
			ClientID:     OAuth2.ClientID,
			ClientSecret: OAuth2.ClientSecret,
			TokenURL:     OAuth2.TokenEndepoint,
			Scopes:       []string{},
		}
		tokenSource = config.TokenSource(ctx)
	}
	return tokenSource
}
