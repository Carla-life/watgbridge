package main

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/akshettrj/watgbridge/pkg/config"
	"github.com/akshettrj/watgbridge/pkg/telegram"

	"github.com/spf13/viper"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("error while loading config: %s", err))
	}

	var (
		botToken           = viper.GetString("telegram.bot_token")
		apiURL             = viper.GetString("telegram.api_url")
		dropPendingUpdates = viper.GetBool("telegram.drop_pending_updates")

		linkPreviews = viper.GetBool("whatsapp.bridging.link_previews")
	)

	bot, updater, err := telegram.NewClient(botToken, apiURL, !linkPreviews)
	if err != nil {
		panic(fmt.Sprintf("failed to create telegram client: %s", err))
	}

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: dropPendingUpdates,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 10,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to start telegram polling: %s", err))
	}
}
