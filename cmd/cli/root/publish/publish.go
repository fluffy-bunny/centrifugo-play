package publish

import (
	"centrifugo-play/cmd/cli/internal"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/centrifugal/centrifuge-go"
	zerolog "github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	cc "golang.org/x/oauth2/clientcredentials"
)

// Version global
var Message string

func init() {
}

// Init command
func Init(rootCmd *cobra.Command) {
	var command = &cobra.Command{
		Use:   "publish",
		Short: "publish to a channel",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Timestamp().Logger()

			config := &cc.Config{
				ClientID:     internal.OAuth2.ClientID,
				ClientSecret: internal.OAuth2.ClientSecret,
				TokenURL:     internal.OAuth2.TokenEndepoint,
				Scopes:       []string{},
			}
			tokenSource := config.TokenSource(ctx)
			token, err := tokenSource.Token()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to get token")
			}
			log.Info().Interface("token", token).Msg("got token")
			endpoint := fmt.Sprintf("ws://localhost:%d/connection/websocket", internal.Port)
			client := centrifuge.NewJsonClient(
				endpoint,
				centrifuge.Config{
					GetToken: func(centrifuge.ConnectionTokenEvent) (string, error) {
						token, err := tokenSource.Token()
						if err != nil {
							return "", err
						}
						return token.AccessToken, nil
					},
				},
			)
			defer client.Close()

			client.OnConnecting(func(e centrifuge.ConnectingEvent) {
				log.Printf("Connecting - %d (%s)", e.Code, e.Reason)
			})
			client.OnConnected(func(e centrifuge.ConnectedEvent) {
				log.Printf("Connected with ID %s", e.ClientID)
			})
			client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
				log.Printf("Disconnected: %d (%s)", e.Code, e.Reason)
			})

			client.OnError(func(e centrifuge.ErrorEvent) {
				log.Printf("Error: %s", e.Error.Error())
			})

			client.OnMessage(func(e centrifuge.MessageEvent) {
				log.Printf("Message from server: %s", string(e.Data))
			})
			client.OnSubscribed(func(e centrifuge.ServerSubscribedEvent) {
				log.Printf("Subscribed to server-side channel %s: (was recovering: %v, recovered: %v)", e.Channel, e.WasRecovering, e.Recovered)
			})
			client.OnSubscribing(func(e centrifuge.ServerSubscribingEvent) {
				log.Printf("Subscribing to server-side channel %s", e.Channel)
			})
			client.OnUnsubscribed(func(e centrifuge.ServerUnsubscribedEvent) {
				log.Printf("Unsubscribed from server-side channel %s", e.Channel)
			})

			client.OnPublication(func(e centrifuge.ServerPublicationEvent) {
				log.Printf("Publication from server-side channel %s: %s (offset %d)", e.Channel, e.Data, e.Offset)
			})
			client.OnJoin(func(e centrifuge.ServerJoinEvent) {
				log.Printf("Join to server-side channel %s: %s (%s)", e.Channel, e.User, e.Client)
			})
			client.OnLeave(func(e centrifuge.ServerLeaveEvent) {
				log.Printf("Leave from server-side channel %s: %s (%s)", e.Channel, e.User, e.Client)
			})

			err = client.Connect()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to connect")
			}
			msg := &internal.ChatMessage{
				Input: Message,
			}
			data, _ := json.Marshal(msg)

			_, err = client.Publish(context.Background(), internal.Channel, data)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to publish message")
			}
			log.Info().Msg("published message")
		},
	}
	var flagName = "message"
	message := `{"a":"b"}`
	command.PersistentFlags().StringVar(&Message, flagName, message, fmt.Sprintf("[required] i.e. --%s='%s'", flagName, message))
	viper.BindPFlag(flagName, rootCmd.PersistentFlags().Lookup(flagName))

	rootCmd.AddCommand(command)
}
