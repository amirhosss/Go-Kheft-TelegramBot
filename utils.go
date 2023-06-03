package main

import (
	"fmt"
	"kheft/bot/languages"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func checkMemberShip(b *gotgbot.Bot) filters.Message {
	return func(m *gotgbot.Message) bool {
		chatMember, err := b.GetChatMember(Configs.channelChatId, m.Chat.Id, nil)
		status := chatMember.GetStatus()
		if err != nil {
			panic(err)
		} else if status == "creator" || status == "member" {
			return true
		} else {
			response := fmt.Sprintf(strings.Join(languages.Response.Messages.Default.Response, "\n"),
				m.Chat.FirstName, m.Chat.Id, Configs.channelUsername)
			_, err := m.Reply(b, response, &gotgbot.SendMessageOpts{
				ParseMode: "MarkdownV2",
			})
			if err != nil {
				panic(err)
			}
			// fmt.Println(response)
			return false
		}
	}
}
