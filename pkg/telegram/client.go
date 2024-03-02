package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

type watgbridgeBaseBot struct {
	disableWebPagePreview bool

	locks map[int64](map[int64]*sync.Mutex)

	gotgbot.BotClient
}

func (b watgbridgeBaseBot) RequestWithContext(ctx context.Context, token string, method string, params map[string]string, data map[string]gotgbot.NamedReader, opts *gotgbot.RequestOpts) (json.RawMessage, error) {
	if strings.HasPrefix(method, "send") || method == "copyMessage" {
		params["allow_sending_without_reply"] = "true"
	}

	if strings.HasPrefix(method, "send") || strings.HasPrefix(method, "edit") {
		params["parse_mode"] = "html"

		if b.disableWebPagePreview {
			params["disable_web_page_preview"] = "true"
		}
	}

	if chatId, found := params["chat_id"]; found {
		if chatId, err := strconv.ParseInt(chatId, 10, 64); err == nil {
			if _, found := b.locks[chatId]; !found {
				b.locks[chatId] = make(map[int64]*sync.Mutex)
			}
			if threadId, found := params["message_thread_id"]; found {
				if threadId, err := strconv.ParseInt(threadId, 10, 64); err == nil {
					if _, found := b.locks[chatId][threadId]; !found {
						b.locks[chatId][threadId] = &sync.Mutex{}
					}
					lock := b.locks[chatId][threadId]
					lock.Lock()
					defer lock.Unlock()
				}
			}
		}
	}

	for {
		res, err := b.BotClient.RequestWithContext(ctx, token, method, params, data, opts)
		if err == nil {
			return res, err
		}

		tgErr, ok := err.(*gotgbot.TelegramError)
		if !ok {
			return res, err
		}

		if tgErr.Code == 429 {
			// Rate limit
			log.Printf("[telegram_rate_limit] sleeping for %d seconds\n", tgErr.ResponseParams.RetryAfter)
			time.Sleep(time.Second * time.Duration(tgErr.ResponseParams.RetryAfter))
			continue
		}

		return res, err
	}
}

func NewClient(
	botToken string, apiURL string,
	disableWebPagePreview bool,
) (*gotgbot.Bot, *ext.Updater, error) {
	bot, err := gotgbot.NewBot(botToken, &gotgbot.BotOpts{
		BotClient: &watgbridgeBaseBot{
			disableWebPagePreview: disableWebPagePreview,
			locks:                 make(map[int64]map[int64]*sync.Mutex),
			BotClient: &gotgbot.BaseBotClient{
				Client: http.Client{},
				DefaultRequestOpts: &gotgbot.RequestOpts{
					APIURL:  apiURL,
					Timeout: time.Minute * 5,
				},
			},
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("could not initialize telegram bot: %s", err)
	}

	dispatcher := ext.NewDispatcher(nil)
	updater := ext.NewUpdater(dispatcher, nil)

	return bot, updater, nil
}
