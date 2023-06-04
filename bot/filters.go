package bot

import (
	"log"

	conf "kheft/bot/configs"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

func CheckMembership(b *gotgbot.Bot, state bool) filters.Message {
	// state parameter reverse the return value

	return func(m *gotgbot.Message) bool {
		chatMember, err := b.GetChatMember(conf.Configs.ChannelChatId, m.Chat.Id, nil)
		if err != nil {
			log.Printf("cannot get status: %s", err)
		}
		status := chatMember.GetStatus()
		if (status == "creator" || status == "member") && state {
			return true
		} else if (status == "creator" || status == "member") && !state {
			return false
		} else if !state {
			return true
		}

		return false
	}
}
