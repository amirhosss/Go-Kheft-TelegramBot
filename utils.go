package main

import (
	"fmt"
	"kheft/bot/languages"
	"log"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func checkMemberShip(b *gotgbot.Bot) filters.Message {
	return func(m *gotgbot.Message) bool {
		chatMember, err := b.GetChatMember(Configs.channelChatId, m.Chat.Id, nil)
		status := chatMember.GetStatus()
		if err != nil {
			log.Printf("Cannot get status: %s", err)
		} else if status == "creator" || status == "member" {
			return true
		} else {
			response := fmt.Sprintf(strings.Join(languages.Response.Messages.Default.Response, "\n"),
				m.Chat.FirstName, m.Chat.Id, Configs.channelUsername)

			var keyboards [][]gotgbot.InlineKeyboardButton
			btns := languages.Response.Messages.Default.Btns
			keyboard := make([]gotgbot.InlineKeyboardButton, len(btns))

			for i, data := range btns {
				keyboard[i].Text = data.Text
				keyboard[i].CallbackData = data.Callback
			}
			keyboards = append(keyboards, keyboard[:])
			markup := gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: keyboards,
			}

			_, err := m.Reply(b, response, &gotgbot.SendMessageOpts{
				ParseMode:   "MarkdownV2",
				ReplyMarkup: markup,
			})
			if err != nil {
				log.Printf("Failed to send reply: %s", err)
			}
		}
		return false
	}
}
