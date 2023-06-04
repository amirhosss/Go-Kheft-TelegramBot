package main

import (
	"log"

	conf "kheft/bot/configs"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func checkMembership(b *gotgbot.Bot, state bool) filters.Message {
	// state parameter reverse the return value

	return func(m *gotgbot.Message) bool {
		chatMember, err := b.GetChatMember(conf.Configs.ChannelChatId, m.Chat.Id, nil)
		if err != nil {
			log.Printf("cannot get status: %s", err)
		}
		status := chatMember.GetStatus()
		if (status == "creator" || status == "member") && state == true {
			return true
		} else if (status == "creator" || status == "member") && state == false {
			return false
		} else if state == false {
			return true
		}

		return false
	}
}
