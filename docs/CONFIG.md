# WaTgBridge Configuration Options

This file has been auto-generated. Do not edit while contributing.

## Config File

### File Formats

The project supports a variety of configuration file formats including **JSON, TOML, YAML, HCL, envfile and Java properties** (all thanks to [spf13/viper](https://github.com/spf13/viper)).

### File Naming

You can create a config file with the name `config` and the extension of your own choice (for file formats listed above). For example, to use JSON format, you will have to name the config file `config.json` and place it in the same directory from which you will execute your binary.

**Make sure that there is only one of such config files at a time in the directory when running the bot as the project will not be able to figure out which file to read and update when changes in config are made through the bot.**

### Environment Variables

You can also use environment variables to configure this bot.

For example, to set the value `bot_api` under `telegram` section, you can use the `WATGBRIDGE_telegram.bot_api` environment variable (i.e. add a `WATGBRIDGE_` prefix).

*The `WATGBRIDGE` part in the environment variable name is case-sensitive, rest of the name is case-insensitive.*

## Config Structure

**Please read the full descriptions carefully.**

The project looks for the following configuration keys:

| Hierarchy | Name | Description | Type | Required | Default Value |
|:---------:|:----:|-------------|:----:|:--------:|:-------------:|
| `-` | `database` | This option is used to configure the database used by the project to store various settings and message IDs. | `config.DatabaseConfig` | Yes | `-` |
| `database > mysql` | `dbname` | The database name to connect with | `string` | No | `-` |
| `database > mysql` | `host` | The hostname/IP address for the database | `string` | No | `-` |
| `database > mysql` | `password` | The password for logging in into the database | `string` | No | `-` |
| `database > mysql` | `port` | The port on which the database is running | `uint64` | No | `-` |
| `database > mysql` | `user` | The username for logging in into the database | `string` | No | `-` |
| `database > postgresql` | `dbname` | The database name to connect with | `string` | No | `-` |
| `database > postgresql` | `disable_ssl` | Whether to disable SSL certificate checking | `bool` | No | `-` |
| `database > postgresql` | `host` | The hostname/IP address for the database | `string` | No | `-` |
| `database > postgresql` | `password` | The password for logging in into the database | `string` | No | `-` |
| `database > postgresql` | `port` | The port on which the database is running | `uint64` | No | `-` |
| `database > postgresql` | `user` | The username for logging in into the database | `string` | No | `-` |
| `database > sqlite` | `path` | The path to the local database file | `string` | No | `-` |
| `telegram` | `api_url` | The telegram bot API URL that you want to use.<br><br>**Note**: The official bot API has restrictions on downloading (\~20MB) and uploading (\~50MB) files. To bypass these limits, you can host your own bot API server the details for which can be found at [telegram-bot-api](https://github.com/tdlib/telegram-bot-api). **Make sure the bot API is hosted at the same server/machine as your bot.** | `string` | No | `https://api.telegram.org/` |
| `telegram` | `bot_token` | The Telegram provided bot token for your bot account which will be used to send messages on Telegram.<br><br>You can get it from [@BotFather](https://t.me/BotFather) | `string` | Yes | `-` |
| `telegram` | `owner_id` | This is the telegram user ID who can control the bot and the messages from whom will be forwarded to WhatsApp. Also, some errors will also be sent to the DM of this account.<br><br>This value can be obtained by sending `/id` to [@MetaButlerBot](https://t.me/Metabutlerbot) or [@MissRose_bot](https://t.me/MissRose_bot).<br><br>**Note**: Make sure to start the bot at least once in its DM by this account because otherwise bot will not be able to send error messages to your DM. | `int64` | Yes | `-` |
| `telegram` | `send_online_presence` | In its default behaviour, the account appears to be ghost and the last seen time is not updated on sending messages through Telegram. You can set this option to `true` so that the account shows up as online for 5 seconds on sending new messages. | `bool` | No | `false` |
| `telegram` | `send_read_receipts` | In its default behaviour, the account appears to be ghost and the messages are not marked as read (blue ticks) even if you send new messages from Telegram to that chat. You can set this option to `true` so that all the pending messages are marked as read upon sending new messages to that chat. | `bool` | No | `false` |
| `telegram` | `sudo_users_id` | This is to mark some other user IDs that can control the bot and send messages through it. This is helpful if you have multiple accounts. | `[]int64` | No | `[]` |
| `telegram` | `update_commands` | This option can be set to `false` so that the commands list which appears when you type `/` in a chat with bots is not updated with commands from this project. | `bool` | No | `true` |
| `updates` | `check_for_updates` | Setting this to `false` will stop the bot from checking for project updates automatically. | `bool` | No | `true` |
| `updates` | `only_releases` | If set to `false` then the bot will also notify if there are new commits on GitHub but a new version release has not been made. | `bool` | No | `true` |
| `whatsapp > bridging` | `audios` | Setting this to `false` will stop bridging audios from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `contacts` | Setting this to `false` will stop bridging contacts from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `documents` | Setting this to `false` will stop bridging documents from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `gifs` | Setting this to `false` will stop bridging gifs from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `images` | Setting this to `false` will stop bridging images from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `large_text_as_files` | If set to `true`, large messages will be sent in a text file. Otherwise, they will be cut short to fit the Telegram's message length limit. | `bool` | No | `true` |
| `whatsapp > bridging` | `link_previews` | Setting this to `false` will disable all link previews in messages forwarded to Telegram by the bot. | `bool` | No | `true` |
| `whatsapp > bridging` | `locations` | Setting this to `false` will stop bridging locations from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `status` | Setting this to `false` will stop bridging statuses from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `stickers` | Setting this to `false` will stop bridging statuses from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `video_stickers` | Setting this to `false` will stop bridging video stickers from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `videos` | Setting this to `false` will stop bridging videos from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp > bridging` | `voice_notes` | Setting this to `false` will stop bridging voice notes from WhatsApp to Telegram. | `bool` | No | `true` |
| `whatsapp` | `login_database` | This option is used to configure the database used by whatsmeow library to store WhatsApp login session. | `config.DatabaseConfig` | Yes | `-` |
| `whatsapp > login_database > mysql` | `dbname` | The database name to connect with | `string` | No | `-` |
| `whatsapp > login_database > mysql` | `host` | The hostname/IP address for the database | `string` | No | `-` |
| `whatsapp > login_database > mysql` | `password` | The password for logging in into the database | `string` | No | `-` |
| `whatsapp > login_database > mysql` | `port` | The port on which the database is running | `uint64` | No | `-` |
| `whatsapp > login_database > mysql` | `user` | The username for logging in into the database | `string` | No | `-` |
| `whatsapp > login_database > postgresql` | `dbname` | The database name to connect with | `string` | No | `-` |
| `whatsapp > login_database > postgresql` | `disable_ssl` | Whether to disable SSL certificate checking | `bool` | No | `-` |
| `whatsapp > login_database > postgresql` | `host` | The hostname/IP address for the database | `string` | No | `-` |
| `whatsapp > login_database > postgresql` | `password` | The password for logging in into the database | `string` | No | `-` |
| `whatsapp > login_database > postgresql` | `port` | The port on which the database is running | `uint64` | No | `-` |
| `whatsapp > login_database > postgresql` | `user` | The username for logging in into the database | `string` | No | `-` |
| `whatsapp > login_database > sqlite` | `path` | The path to the local database file | `string` | No | `-` |
| `whatsapp` | `session_name` | This session name will show up when you open the list of `Linked Devices` in your phone's WhatsApp mobile application. | `string` | No | `WaTgBridge` |
| `whatsapp > sticker_metadata` | `author_name` | This is the pack author name which will be set for the stickers sent to WhatsApp from Telegram. | `string` | No | `WaTgBridge` |
| `whatsapp > sticker_metadata` | `pack_name` | This is the pack name which will be set for the stickers sent to WhatsApp from Telegram. | `string` | No | `WaTgBridge` |
