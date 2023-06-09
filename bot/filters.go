package bot

import (
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"
)

type CheckMembershipOpts struct {
	// state parameter should reverse the return value
	ReverseState bool
	MessageText  string
}

func (opts *CheckMembershipOpts) checkText(m *gotgbot.Message) bool {
	if opts.MessageText == "" || opts.MessageText == m.Text {
		return true
	}
	return false
}

// This function could check the membership of specific channel and equality of text message
func (opts *CheckMembershipOpts) CheckMessage(b *gotgbot.Bot) filters.Message {

	return func(m *gotgbot.Message) bool {
		chatMember, err := b.GetChatMember(Configs.ChannelChatId, m.Chat.Id, nil)
		if err != nil {
			log.Printf("cannot get status: %s", err)
		}

		switch status := chatMember.GetStatus(); status {
		case "creator", "member":
			if opts != nil {
				if !opts.ReverseState && opts.checkText(m) {
					return true
				} else {
					return false
				}
			}
			return true
		default:
			if opts != nil {
				if opts.ReverseState && opts.checkText(m) {
					return true
				}
			}
		}
		return false
	}
}
