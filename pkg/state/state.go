package state

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"go.mau.fi/whatsmeow"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const WATGBRIDGE_VERSION = "2.0.0"

type state struct {
	Database *gorm.DB
	Logger   *zap.Logger

	Telegram struct {
		Bot        *gotgbot.Bot
		Dispatcher *ext.Dispatcher
		Updater    *ext.Updater
		Commands   []gotgbot.BotCommand
	}

	WhatsAppClient *whatsmeow.Client

	StartTime     time.Time
	LocalLocation *time.Location
}

var State state
