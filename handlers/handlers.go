package handlers

import (
	"fmt"
	"strings"

	conf "kheft/bot/configs"
	"kheft/bot/languages"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// echo replies to a messages with its own contents.
func NonMemberStart(b *gotgbot.Bot, ctx *ext.Context) error {
	response := fmt.Sprintf(strings.Join(languages.Response.Messages.Default.Response, "\n"),
		ctx.Message.Chat.FirstName, ctx.Message.Chat.Id, conf.Configs.ChannelUsername)

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

	_, err := ctx.EffectiveMessage.Reply(b, response, &gotgbot.SendMessageOpts{
		ParseMode:   "MarkdownV2",
		ReplyMarkup: markup,
	})
	if err != nil {
		return fmt.Errorf("failed to send reply: %w", err)
	}
	return nil
}
