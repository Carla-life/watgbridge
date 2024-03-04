package config

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

type ConfigOption struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Required    bool            `json:"required"`
	Mutable     bool            `json:"mutable"`
	Default     any             `json:"default"`
	Type        reflect.Type    `json:"type"`
	Hierarchy   []string        `json:"hierarchy"`
	Validator   OptionValidator `json:"-"`
}

func (opt ConfigOption) ViperKey() string {
	if len(opt.Hierarchy) == 0 {
		return opt.Name
	}

	return strings.Join(opt.Hierarchy, ".") + "." + opt.Name
}

func genDatabaseSchema(hierarchy []string) []ConfigOption {
	return []ConfigOption{
		// SQLite
		{
			Name:        "path",
			Description: "The path to the local database file",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "sqlite"),
			Validator:   NonEmptyStringValidator,
		},
		// PostgreSQL
		{
			Name:        "user",
			Description: "The username for logging in into the database",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "postgresql"),
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "password",
			Description: "The password for logging in into the database",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "postgresql"),
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "host",
			Description: "The hostname/IP address for the database",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "postgresql"),
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "port",
			Description: "The port on which the database is running",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(uint64(0)),
			Hierarchy:   append(hierarchy, "postgresql"),
			Validator:   NonZeroNumberValidator[uint64],
		},
		{
			Name:        "dbname",
			Description: "The database name to connect with",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "postgresql"),
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "disable_ssl",
			Description: "Whether to disable SSL certificate checking",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(false),
			Hierarchy:   append(hierarchy, "postgresql"),
			Validator:   nil,
		},
		// MySQL
		{
			Name:        "user",
			Description: "The username for logging in into the database",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "mysql"),
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "password",
			Description: "The password for logging in into the database",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "mysql"),
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "host",
			Description: "The hostname/IP address for the database",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "mysql"),
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "port",
			Description: "The port on which the database is running",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(uint64(0)),
			Hierarchy:   append(hierarchy, "mysql"),
			Validator:   NonZeroNumberValidator[uint64],
		},
		{
			Name:        "dbname",
			Description: "The database name to connect with",
			Required:    false,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   append(hierarchy, "mysql"),
			Validator:   NonEmptyStringValidator,
		},
	}
}

var AllConfigOptions = slices.Concat(
	[]ConfigOption{
		{
			Name:        "git",
			Description: "The path to the git executable",
			Required:    false,
			Mutable:     true,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"executables"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "go",
			Description: "The path to the Golang executable",
			Required:    false,
			Mutable:     true,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"executables"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "ffmpeg",
			Description: "The path to the FFmpeg executable",
			Required:    false,
			Mutable:     true,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"executables"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "bot_token",
			Description: "The Telegram provided bot token for your bot account which will be used to send messages on Telegram.<br><br>You can get it from [@BotFather](https://t.me/BotFather)",
			Required:    true,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"telegram"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "api_url",
			Description: "The telegram bot API URL that you want to use.<br><br>**Note**: The official bot API has restrictions on downloading (\\~20MB) and uploading (\\~50MB) files. To bypass these limits, you can host your own bot API server the details for which can be found at [telegram-bot-api](https://github.com/tdlib/telegram-bot-api). **Make sure the bot API is hosted at the same server/machine as your bot.**",
			Required:    false,
			Mutable:     false,
			Default:     "https://api.telegram.org/",
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"telegram"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "drop_pending_updates",
			Description: "Whether the pending updates from Telegram should be dropped.",
			Required:    true,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"telegram"},
			Validator:   nil,
		},
		{
			Name:        "owner_id",
			Description: "This is the telegram user ID who can control the bot and the messages from whom will be forwarded to WhatsApp. Also, some errors will also be sent to the DM of this account.<br><br>This value can be obtained by sending `/id` to [@MetaButlerBot](https://t.me/Metabutlerbot) or [@MissRose_bot](https://t.me/MissRose_bot).<br><br>**Note**: Make sure to start the bot at least once in its DM by this account because otherwise bot will not be able to send error messages to your DM.",
			Required:    true,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(int64(0)),
			Hierarchy:   []string{"telegram"},
			Validator:   NonZeroNumberValidator[int64],
		},
		{
			Name:        "sudo_users_id",
			Description: "This is to mark some other user IDs that can control the bot and send messages through it. This is helpful if you have multiple accounts.",
			Required:    false,
			Mutable:     true,
			Default:     []int64{},
			Type:        reflect.TypeOf([]int64{}),
			Hierarchy:   []string{"telegram"},
			Validator:   nil,
		},
		{
			Name:        "update_commands",
			Description: "This option can be set to `false` so that the commands list which appears when you type `/` in a chat with bots is not updated with commands from this project.",
			Required:    false,
			Mutable:     false,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"telegram"},
			Validator:   nil,
		},
		{
			Name:        "send_online_presence",
			Description: "In its default behaviour, the account appears to be ghost and the last seen time is not updated on sending messages through Telegram. You can set this option to `true` so that the account shows up as online for 5 seconds on sending new messages.",
			Required:    false,
			Mutable:     true,
			Default:     false,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"telegram"},
			Validator:   nil,
		},
		{
			Name:        "send_read_receipts",
			Description: "In its default behaviour, the account appears to be ghost and the messages are not marked as read (blue ticks) even if you send new messages from Telegram to that chat. You can set this option to `true` so that all the pending messages are marked as read upon sending new messages to that chat.",
			Required:    false,
			Mutable:     true,
			Default:     false,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"telegram"},
			Validator:   nil,
		},
		{
			Name:        "session_name",
			Description: "This session name will show up when you open the list of `Linked Devices` in your phone's WhatsApp mobile application.",
			Required:    false,
			Mutable:     false,
			Default:     "WaTgBridge",
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"whatsapp"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "login_database",
			Description: "This option is used to configure the database used by whatsmeow library to store WhatsApp login session.",
			Required:    true,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(DatabaseConfig{}),
			Hierarchy:   []string{"whatsapp"},
			Validator:   DatabaseValidator,
		},
		{
			Name:        "documents",
			Description: "Setting this to `false` will stop bridging documents from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "images",
			Description: "Setting this to `false` will stop bridging images from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "gifs",
			Description: "Setting this to `false` will stop bridging gifs from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "videos",
			Description: "Setting this to `false` will stop bridging videos from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "voice_notes",
			Description: "Setting this to `false` will stop bridging voice notes from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "audios",
			Description: "Setting this to `false` will stop bridging audios from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "status",
			Description: "Setting this to `false` will stop bridging statuses from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "stickers",
			Description: "Setting this to `false` will stop bridging statuses from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "video_stickers",
			Description: "Setting this to `false` will stop bridging video stickers from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "contacts",
			Description: "Setting this to `false` will stop bridging contacts from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "locations",
			Description: "Setting this to `false` will stop bridging locations from WhatsApp to Telegram.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "link_previews",
			Description: "Setting this to `false` will disable all link previews in messages forwarded to Telegram by the bot.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "large_text_as_files",
			Description: "If set to `true`, large messages will be sent in a text file. Otherwise, they will be cut short to fit the Telegram's message length limit.",
			Required:    false,
			Mutable:     true,
			Default:     true,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"whatsapp", "bridging"},
			Validator:   nil,
		},
		{
			Name:        "pack_name",
			Description: "This is the pack name which will be set for the stickers sent to WhatsApp from Telegram. (This requires the webptool `webpmux` to be installed)",
			Required:    false,
			Mutable:     true,
			Default:     "WaTgBridge",
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"whatsapp", "sticker_metadata"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "author_name",
			Description: "This is the pack author name which will be set for the stickers sent to WhatsApp from Telegram. (This requires the webptool `webpmux` to be installed)",
			Required:    false,
			Mutable:     true,
			Default:     "WaTgBridge",
			Type:        reflect.TypeOf(""),
			Hierarchy:   []string{"whatsapp", "sticker_metadata"},
			Validator:   NonEmptyStringValidator,
		},
		{
			Name:        "check_for_updates",
			Description: "Setting this to `false` will stop the bot from checking for project updates automatically.",
			Required:    true,
			Mutable:     true,
			Default:     nil,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"updates"},
			Validator:   nil,
		},
		{
			Name:        "only_releases",
			Description: "If set to `false` then the bot will also notify if there are new commits on GitHub but a new version release (with pre-built binary) has not been made.",
			Required:    true,
			Mutable:     true,
			Default:     nil,
			Type:        reflect.TypeOf(false),
			Hierarchy:   []string{"updates"},
			Validator:   nil,
		},
		{
			Name:        "database",
			Description: "This option is used to configure the database used by the project to store various settings and message IDs.",
			Required:    true,
			Mutable:     false,
			Default:     nil,
			Type:        reflect.TypeOf(DatabaseConfig{}),
			Hierarchy:   []string{},
			Validator:   DatabaseValidator,
		},
	},
	genDatabaseSchema([]string{"whatsapp", "login_database"}),
	genDatabaseSchema([]string{"database"}),
)

var MutableConfigOptions map[string]interface{}

func init() {
	// Validate the schema while development
	for optNum, opt := range AllConfigOptions {
		if opt.Name == "" {
			panic(fmt.Sprintf("Option number %v found without name", optNum))
		}
		if opt.Description == "" {
			panic(fmt.Sprintf("Option '%s' found without description", opt.ViperKey()))
		}
		if opt.Type == nil {
			panic(fmt.Sprintf("Option '%s' found without type", opt.ViperKey()))
		}
		if opt.Default != nil && reflect.TypeOf(opt.Default) != opt.Type {
			panic(fmt.Sprintf("Option '%s' default is not of the same type as the defined type", opt.ViperKey()))
		}
		if opt.Default != nil && opt.Validator != nil {
			if err := opt.Validator(opt, opt.Default); err != nil {
				panic(fmt.Sprintf("Option '%s' default value does not passes through its validator: %s", opt.ViperKey(), err))
			}
		}
	}

	// Create a hierarchy of mutable options
	MutableConfigOptions = make(map[string]interface{})

	for _, opt := range AllConfigOptions {
		if !opt.Mutable {
			continue
		}
		opt := opt

		if len(opt.Hierarchy) == 0 {
			MutableConfigOptions[opt.Name] = opt
			continue
		}

		currMap := MutableConfigOptions

		for _, hierarchy := range opt.Hierarchy {
			hierarchy := hierarchy

			val, found := currMap[hierarchy]
			if !found {
				currMap[hierarchy] = make(map[string]interface{})
				currMap = currMap[hierarchy].(map[string]interface{})
			} else {
				currMap = val.(map[string]interface{})
			}
		}

		currMap[opt.Name] = opt
	}
}
