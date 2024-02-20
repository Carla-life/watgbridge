# WaTgBridge Configuration Options

## Config File

### File Formats

The project supports a variety of configration file formats including **JSON, TOML, YAML, HCL, envfile and Java properties** (all thanks to [spf13/viper](https://github.com/spf13/viper)).

### File Naming

You can create a config file with the name `config` and the extension of your own choice (from above listed formats). For example, to use JSON format, you will have to name the config file `config.json` and place it in the same directory from which you will execute your binary.

**Make sure that there is only one of such config files at a time in the directory when running the bot as the project will not be able to figure out which file to read and update when changes in config are made through the bot.**

### Environment Variables

You can also use environment variables to configure this bot

For example, to set the value `bot_api` under `telegram` section, you can use the `WATGBRIDGE_telegram.bot_api` environment variable (i.e. add a `WATGBRIDGE_` prefix).

*The `WATGBRIDGE_` part in the environment variable name is case-sensitive, rest of the name is case-insensitive.*

## Config Structure

**Please read the full descriptions carefully.**

The project looks for the following configuration keys:

### Section telegram

These settings will configure the Telegram side of this project

| Hierarchy | Name | Description | Required | Type | Default Value |
|:---------:|:----:|-------------|:--------:|:----:|:-------------:|
| `telegram` | `bot_token` | The telegram provided bot token for you bot account which will be used to send messages on Telegram.<br><br>You can get it by messaging [@BotFather](https://t.me/BotFather) | Yes | string | - |
| `telegram` | `api_url` | The telegram bot API URL that you want to use.<br><br>**Note**: The official bot API has restrictions on downloading (\~20MB) and uploading (\~50MB) files. To bypass these limits, you can host your own bot API server the details for which can be found at [telegram-bot-api](https://github.com/tdlib/telegram-bot-api). **Make sure the bot API is hosted at the same server/machine as your bot.** | No | string | https://api.telegram.org/ |
| `telegram` | `owner_id` | This is the Telegram user ID who can control the bot and the messages from whom will be forwarded to WhatsApp. Also, some errors will also be sent to the DM of this account.<br><br>This value can be obtained by sending `/id` to [@MetaButlerBot](https://t.me/Metabutlerbot) or [@MissRose_bot](https://t.me/MissRose_bot).<br><br>**Note**: Make sure to start the bot atleast once in its DM by this account because otherwise bot will not be able to send error messages to your DM. | Yes | integer | - |
| `telegram` | `sudo_users_id` | This is to mark some other user IDs that can control the bot and send messages through it. This is helpful if you have multiple accounts | No | list of integer | - |
| `telegram` | `update_commands` | This option can be set to `false` so that the commands list which appears when you type `/` in a chat with bots is not updated with commands from this project | No | boolean | true |
| `telegram` | `send_online_presence` | In its default behaviour, the account appears to be a ghost and the last seen time is not updated on sending messages through Telegram. You can set this option to `true` so that the account shows up as online for 5 seconds on sending messages. | No | boolean | false |
| `telegram` | `send_read_receipts` | In its default behaviour, the old messages sent by others are not marked as read even when you send new messages in the chat. You can set this option to `true` so that the old messages are marked as read when you send new messages in a chat through Telegram. | No | boolean | false |

### Section whatsapp

These settings will configure the WhatsApp side of the project

| Hierarchy | Name | Description | Required | Type | Default Value |
|:---------:|:----:|-------------|:--------:|:----:|:-------------:|
| `whatsapp` | `session_name` | This session name will show up when you open the list of `Linked Devices` in your phone's WhatsApp application | No | string | WaTgBridge |
| `whatsapp` => `login_database` | `type` | This the type of SQL database that will be used to store the WhatsApp login session. It can be one of `"sqlite", "postgresql", "mysql"` | Yes | string | - |
| `whatsapp` => `login_database` | `settings` | Please refer [database section](#section-database) for more | Yes | - | - |
| `whatsapp` => `bridging` | `documents` | Setting this to `false` will stop bridging documents from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `images` | Setting this to `false` will stop bridging images from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `gifs` | Setting this to `false` will stop bridging GIFs from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `videos` | Setting this to `false` will stop bridging videos from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `voice_notes` | Setting this to `false` will stop bridging voice notes from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `audios` | Setting this to `false` will stop bridging audios from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `status` | Setting this to `false` will stop bridging status updates from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `stickers` | Setting this to `false` will stop bridging stickers from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `video_stickers` | Setting this to `false` will stop bridging video stickers from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `contacts` | Setting this to `false` will stop bridging contacts from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `locations` | Setting this to `false` will stop bridging location updates from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `link_previews` | Setting this to `false` will disable link previews in updates from WhatsApp to Telegram | No | boolean | true |
| `whatsapp` => `bridging` | `large_text_as_files` | Setting this to `false` will stop sending large texts (that exceed Telegram message length limits) as text files | No | boolean | true |
| `whatsapp` => `sticker_metadata` | `pack_name` | This is the pack name which will be set for the stickers sent to WhatsApp from Telegram | No | string | WaTgBridge |
| `whatsapp` => `sticker_metadata` | `author_name` | This is the pack's author name which will be set for the stickers sent to WhatsApp from Telegram | No | string | WaTgBridge |

### Section database

These settings will configure settings related to the project Database.

| Hierarchy | Name | Description | Required | Type | Default Value |
|:---------:|:----:|-------------|:--------:|:----:|:-------------:|
| `database` | `type` | 
