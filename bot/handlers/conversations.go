package handlers

import (
	"fmt"
	"strings"

	"kheft/bot"
	"kheft/bot/languages"

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
	responseField := languages.Response.Conversations.Registration
	var keyboards [][]gotgbot.KeyboardButton
	btns := responseField.Btns
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
		strings.Join(responseField.Response, "\n"),
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
	responseField := languages.Response.Conversations.Rules
	if ctx.EffectiveMessage.Text == responseField.Query {
		_, err := ctx.EffectiveMessage.Reply(b,
			strings.Join(responseField.Response, "\n"),
			&gotgbot.SendMessageOpts{
				ParseMode: "MarkdownV2",
			},
		)
		if err != nil {
			return fmt.Errorf("rules acceptance failed: %s", err)
		}
		return handlers.NextConversationState("username")
	} else {
		_, err := ctx.EffectiveMessage.Reply(b, responseField.Failed,
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
	responseField := languages.Response.Conversations.Username
	_, err := ctx.EffectiveMessage.Reply(b,
		strings.Join(responseField.Response, "\n"),
		&gotgbot.SendMessageOpts{
			ParseMode: "MARKDOWNV2",
		},
	)
	if err != nil {
		return fmt.Errorf("get username failed: %s", err)
	}
	return handlers.NextConversationState("price")
}

func GetPrice(b *gotgbot.Bot, ctx *ext.Context) error {
	responseField := languages.Response.Conversations.Price
	if strings.HasPrefix(ctx.EffectiveMessage.Text, "@") {
		_, err := ctx.EffectiveMessage.Reply(b,
			strings.Join(responseField.Response, "\n"),
			&gotgbot.SendMessageOpts{
				ParseMode: "MarkdownV2",
			},
		)
		if err != nil {
			return fmt.Errorf("get price failed: %s", err)
		}
		return handlers.EndConversation()
	} else {
		_, err := ctx.EffectiveMessage.Reply(b,
			responseField.Failed,
			&gotgbot.SendMessageOpts{
				ParseMode: "MARKDOWNV2",
			},
		)
		if err != nil {
			return fmt.Errorf("get username failed: %s", err)
		}
	}
	return handlers.NextConversationState("price")
}
