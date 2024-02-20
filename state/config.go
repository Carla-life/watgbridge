package state

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Type string `mapstructure:"type"`

	Settings struct {
		Sqlite *struct {
			Path string `mapstructure:"path"`
		} `mapstructure:"sqlite"`

		PostgresQL *struct {
			User       string `mapstructure:"user"`
			Password   string `mapstructure:"password"`
			Host       string `mapstructure:"host"`
			Port       int64  `mapstructure:"port"`
			DBName     string `mapstructure:"dbname"`
			DisableSSL bool   `mapstructure:"disable_ssl"`
		} `mapstructure:"postgresql"`

		MySQL *struct {
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Host     string `mapstructure:"host"`
			Port     int64  `mapstructure:"port"`
			DBName   string `mapstructure:"dbname"`
		} `mapstructure:"mysql"`
	} `mapstructure:"settings"`
}

type WaTgBridgeConfig struct {
	Telegram struct {
		BotToken            string  `mapstructure:"bot_token"`
		APIURL              string  `mapstructure:"api_url"`
		OwnerId             int64   `mapstructure:"owner_id"`
		SudoUsersID         []int64 `mapstructure:"sudo_users_id"`
		UpdateCommands      bool    `mapstructure:"update_commands"`
		SilentConfirmations bool    `mapstructure:"silent_confirmations"`
		SendOnlinePresence  bool    `mapstructure:"send_online_presence"`
		SendReadReceipts    bool    `mapstructure:"send_read_receipts"`
	} `mapstructure:"telegram"`
}

func LoadConfig() error {
	// Read config from the current directory.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("WATGBRIDGE")
	viper.AutomaticEnv()

	// Watch the config files for changes and reload.
	viper.OnConfigChange(func(e fsnotify.Event) {
		// log.Printf("Config file changed")
	})
	viper.WatchConfig()

	return viper.ReadInConfig()
}
