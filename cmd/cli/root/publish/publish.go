package publish

import (
	"centrifugo-play/cmd/cli/internal"
	"context"
	"encoding/json"
	"fmt"
	"os"

	centrifuge "github.com/centrifugal/centrifuge-go"
	zerolog "github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)

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

			tokenSource := internal.GetTokenSource(ctx)
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
				log.Info().Interface("event", e).Msg("OnConnecting")
			})
			client.OnConnected(func(e centrifuge.ConnectedEvent) {
				log.Info().Interface("event", e).Msg("OnConnected")
			})
			client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
				log.Info().Interface("event", e).Msg("OnDisconnected")
			})
			client.OnError(func(e centrifuge.ErrorEvent) {
				log.Error().Interface("event", e).Msg("OnError")
			})

			client.OnMessage(func(e centrifuge.MessageEvent) {
				log.Info().Interface("event", e).Msg("OnMessage")
			})
			client.OnSubscribed(func(e centrifuge.ServerSubscribedEvent) {
				log.Info().Interface("event", e).Msg("OnSubscribed")
			})
			client.OnSubscribing(func(e centrifuge.ServerSubscribingEvent) {
				log.Info().Interface("event", e).Msg("OnSubscribing")
			})
			client.OnUnsubscribed(func(e centrifuge.ServerUnsubscribedEvent) {
				log.Info().Interface("event", e).Msg("OnUnsubscribed")
			})

			client.OnPublication(func(e centrifuge.ServerPublicationEvent) {
				log.Info().Interface("event", e).Msg("OnPublication")
			})
			client.OnJoin(func(e centrifuge.ServerJoinEvent) {
				log.Info().Interface("event", e).Msg("OnJoin")
			})
			client.OnLeave(func(e centrifuge.ServerLeaveEvent) {
				log.Info().Interface("event", e).Msg("OnLeave")
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
	viper.BindPFlag(flagName, command.PersistentFlags().Lookup(flagName))

	rootCmd.AddCommand(command)
}
