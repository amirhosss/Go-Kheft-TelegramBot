package handlers

import (
	"fmt"
	"kheft/bot"
	"kheft/bot/languages"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Exit(b *gotgbot.Bot, ctx *ext.Context) error {
	err := MemberStart(b, ctx)
	if err != nil {
		return fmt.Errorf("exit conversation failed: %s", err)
	}
	return handlers.EndConversation()
}

func Registration(b *gotgbot.Bot, ctx *ext.Context) error {
	var keyboards [][]gotgbot.KeyboardButton
	btns := languages.Response.Conversations.Registration.Btns
	keyboard := make([]gotgbot.KeyboardButton, len(btns))

	for i, data := range btns {
		keyboard[i].Text = data
	}
	keyboards = append(keyboards, keyboard[:])
	markup := gotgbot.ReplyKeyboardMarkup{
		Keyboard:        keyboards,
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	response := fmt.Sprintf(
		strings.Join(languages.Response.Conversations.Registration.Response, "\n"),
		bot.Configs.RegistrationPrice,
	)
	_, err := ctx.EffectiveMessage.Reply(b,
		response,
		&gotgbot.SendMessageOpts{
			ReplyMarkup: markup,
			ParseMode:   "MarkdownV2",
		},
	)
	if err != nil {
		return fmt.Errorf("registration failed: %s", err)
	}
	return handlers.NextConversationState("rules")
}

func RulesAcceptance(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveMessage.Text == languages.Response.Conversations.Registration.Btns[0] {
		_, err := ctx.EffectiveMessage.Reply(b,
			strings.Join(languages.Response.Conversations.Rules.Response, "\n"),
			&gotgbot.SendMessageOpts{
				ParseMode: "MarkdownV2",
			},
		)
		if err != nil {
			return fmt.Errorf("rules acceptance failed: %s", err)
		}
		return handlers.NextConversationState("username")
	} else {
		_, err := ctx.EffectiveMessage.Reply(b, languages.Response.Conversations.Rules.Failed,
			&gotgbot.SendMessageOpts{
				ParseMode: "MARKDOWNV2",
			})
		if err != nil {
			return fmt.Errorf("rules acceptance failed: %s", err)

		}
	}
	return handlers.NextConversationState("rules")
}

func GetUsername(b *gotgbot.Bot, ctx *ext.Context) error {
	if strings.HasPrefix(ctx.EffectiveMessage.Text, "@") {
		_, err := ctx.EffectiveMessage.Reply(b,
			strings.Join(languages.Response.Conversations.Username.Response, "\n"),
			&gotgbot.SendMessageOpts{
				ParseMode: "MarkdownV2",
			},
		)
		if err != nil {
			return fmt.Errorf("get username failed: %s", err)
		}
		return handlers.NextConversationState("price")
	} else {
		_, err := ctx.EffectiveMessage.Reply(b,
			languages.Response.Conversations.Username.Failed,
			&gotgbot.SendMessageOpts{
				ParseMode: "MARKDOWNV2",
			},
		)
		if err != nil {
			return fmt.Errorf("get username failed: %s", err)
		}
	}

	return handlers.NextConversationState("username")
}
