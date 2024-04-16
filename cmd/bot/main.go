package main

import (
	"fmt"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/akshettrj/watgbridge/cmd/bot/config"
	pkg_config "github.com/akshettrj/watgbridge/pkg/config"
	pkg_telegram "github.com/akshettrj/watgbridge/pkg/telegram"

	"github.com/spf13/viper"
)

func main() {
	err := pkg_config.LoadConfig()
	if err != nil {
		log.Fatalf("error while loading config: %s", err)
	}

	var (
		botToken           = viper.GetString("telegram.bot_token")
		apiURL             = viper.GetString("telegram.api_url")
		dropPendingUpdates = viper.GetBool("telegram.drop_pending_updates")

		linkPreviews = viper.GetBool("whatsapp.bridging.link_previews")
	)

	bot, updater, err := pkg_telegram.NewClient(botToken, apiURL, !linkPreviews)
	if err != nil {
		log.Fatalf("failed to create telegram client: %s", err)
	}

	if err = config.ValidateConfig(bot); err != nil {
		log.Fatalf("problems with your config: %s", err)
	}

	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: dropPendingUpdates,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 10,
		},
	})
	if err != nil {
		log.Fatalf("failed to start telegram polling: %s", err)
	}

	updater.Idle()
}
