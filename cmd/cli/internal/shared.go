package internal

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
)

type ChatMessage struct {
	Input string `json:"input"`
}
