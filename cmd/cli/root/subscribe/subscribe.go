package subscribe

import (
	"centrifugo-play/cmd/cli/internal"
	"context"
	"fmt"
	"os"

	"github.com/centrifugal/centrifuge-go"
	zerolog "github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
)

var WithChannelToken bool

func init() {
}

// Init command
func Init(rootCmd *cobra.Command) {
	var command = &cobra.Command{
		Use:   "subscribe",
		Short: "subscribe to a channel",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Timestamp().Logger()
			log := log.With().Str("channel", internal.Channel).Logger()

			tokenSource := internal.GetTokenSource(ctx)
			token, err := tokenSource.Token()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to get token")
			}
			log.Info().Interface("token", token).Msg("got token")
			clientLog := log.With().Str("context", "client").Logger()
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
				clientLog.Info().Interface("event", e).Msg("OnConnecting")
			})
			client.OnConnected(func(e centrifuge.ConnectedEvent) {
				clientLog.Info().Interface("event", e).Msg("OnConnected")
			})
			client.OnDisconnected(func(e centrifuge.DisconnectedEvent) {
				clientLog.Info().Interface("event", e).Msg("OnDisconnected")
			})
			client.OnError(func(e centrifuge.ErrorEvent) {
				clientLog.Error().Interface("event", e).Msg("OnError")
			})

			client.OnMessage(func(e centrifuge.MessageEvent) {
				clientLog.Info().Interface("event", e).Msg("OnMessage")
			})
			client.OnSubscribed(func(e centrifuge.ServerSubscribedEvent) {
				clientLog.Info().Interface("event", e).Msg("OnSubscribed")
			})
			client.OnSubscribing(func(e centrifuge.ServerSubscribingEvent) {
				clientLog.Info().Interface("event", e).Msg("OnSubscribing")
			})
			client.OnUnsubscribed(func(e centrifuge.ServerUnsubscribedEvent) {
				clientLog.Info().Interface("event", e).Msg("OnUnsubscribed")
			})

			client.OnPublication(func(e centrifuge.ServerPublicationEvent) {
				clientLog.Info().Interface("event", e).Msg("OnPublication")
			})
			client.OnJoin(func(e centrifuge.ServerJoinEvent) {
				clientLog.Info().Interface("event", e).Msg("OnJoin")
			})
			client.OnLeave(func(e centrifuge.ServerLeaveEvent) {
				clientLog.Info().Interface("event", e).Msg("OnLeave")
			})

			err = client.Connect()
			if err != nil {
				clientLog.Fatal().Err(err).Msg("failed to connect")
			}
			subscriptionLog := log.With().Str("context", "subscribe").Logger()

			sub, err := client.NewSubscription(internal.Channel,
				centrifuge.SubscriptionConfig{
					Recoverable: true,
					JoinLeave:   true,
					Positioned:  true,
				})
			if err != nil {
				subscriptionLog.Fatal().Err(err).Msg("failed to create subscription")
			}
			sub.OnSubscribed(func(e centrifuge.SubscribedEvent) {
				subscriptionLog.Info().Interface("event", e).Msg("OnSubscribed")
			})
			sub.OnError(func(e centrifuge.SubscriptionErrorEvent) {
				subscriptionLog.Error().Interface("event", e).Msg("OnError")
			})
			sub.OnUnsubscribed(func(e centrifuge.UnsubscribedEvent) {
				subscriptionLog.Info().Interface("event", e).Msg("OnUnsubscribed")
			})
			sub.OnPublication(func(e centrifuge.PublicationEvent) {
				subscriptionLog.Info().Interface("event", e).Msg("OnPublication")
			})
			if !WithChannelToken {
				// Subscribe on private channel.
				err = sub.Subscribe()
				if err != nil {
					subscriptionLog.Fatal().Err(err).Msg("failed to subscribe")
				}
			}
			subscriptionLog.Info().Msg("sub.Subscribe")

			select {}
		},
	}
	var flagName = "with-channel-token"
	command.PersistentFlags().BoolVar(&WithChannelToken, flagName, false, fmt.Sprintf("i.e. --%s=true  This will not call subscribe and rely on the token claims.", flagName))
	viper.BindPFlag(flagName, command.PersistentFlags().Lookup(flagName))

	rootCmd.AddCommand(command)
}
