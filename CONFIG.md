# WaTgBridge Configuration Options

## Config File

The project supports a variety of configration file formats including **JSON, TOML, YAML, HCL, envfile and Java properties** (all thanks to [spf13/viper](https://github.com/spf13/viper)).

You can create a config file with the name `config` and the extension of your own choice (from above listed formats). For example, to use JSON format, you will have to name the config file `config.json` and place it in the same directory from which you will execute your binary.

## Config Structure

**Please read the full descriptions carefully.**

The project looks for the following configuration keys:

### Section `telegram`

| Name | Description | Required | Type | Default Value |
|:----:|-------------|:--------:|:----:|:-------------:|
| `bot_token` | The telegram provided bot token for you bot account which will be used to send messages on Telegram.<br><br>You can get it by messaging [@BotFather](https://t.me/BotFather) | Yes | string | - |
| `api_url` | The telegram bot API URL that you want to use.<br><br>**Note**: The official bot API has restrictions on downloading (\~20MB) and uploading (\~50MB) files. To bypass these limits, you can host your own bot API server the details for which can be found at [telegram-bot-api](https://github.com/tdlib/telegram-bot-api). **Make sure the bot API is hosted at the same server/machine as your bot.** | No | string | https://api.telegram.org/ |
| `owner_id` | This is the Telegram user ID who can control the bot and the messages from whom will be forwarded to WhatsApp. Also, some errors will also be sent to the DM of this account.<br><br>This value can be obtained by sending `/id` to [@MetaButlerBot](https://t.me/Metabutlerbot) or [@MissRose_bot](https://t.me/MissRose_bot).<br><br>**Note**: Make sure to start the bot atleast once in its DM by this account because otherwise bot will not be able to send error messages to your DM. | Yes | integers | - |
| `sudo_users_id` | This is to mark some other user IDs that can control the bot and send messages through it. This is helpful if you have multiple accounts | No | list of integers | - |
| `skip_setting_commands` | This option can be set to `true` so that the commands list which appears when you type `/` in a chat with bots is not updated with commands from this project | No | boolean | false |
| `silent_confirmations` | When sending messages from Telegram, the bot sends a temporary confirmation message. This option can be set to `true` for sending silent notifications for that confirmation message | No | boolean | false |
