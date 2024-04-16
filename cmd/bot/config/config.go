package config

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/spf13/viper"
)

func ValidateConfig(bot *gotgbot.Bot) error {
	var (
		chatId      = viper.GetInt64("telegram.target_chat_id")
		ownerId     = viper.GetInt64("telegram.owner_id")
		sudoUsersId = viper.Get("telegram.sudo_users_id").([]int64)
	)

	chat, err := bot.GetChat(chatId, nil)
	if err != nil {
		return fmt.Errorf("failed to get information about the group: %s", err)
	}

	if !chat.IsForum {
		return fmt.Errorf("topics are not enabled in the target group, go to settings and enable them")
	}

	return nil
}
