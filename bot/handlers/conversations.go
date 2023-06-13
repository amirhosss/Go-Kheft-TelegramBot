package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"kheft/bot"
	"kheft/bot/languages"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var users map[int64]*User

func init() {
	users = make(map[int64]*User)
}

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

	printer := message.NewPrinter(language.Persian)
	response := printer.Sprintf(
		strings.Join(responseField.Response, strings.Repeat("\n", 2)),
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
	users[ctx.EffectiveMessage.Chat.Id] = &User{
		AdvertiseDescription: ctx.EffectiveMessage.Text,
	}
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
		if user, ok := users[ctx.EffectiveMessage.Chat.Id]; ok {
			user.Username = ctx.EffectiveMessage.Text
		}

		_, err := ctx.EffectiveMessage.Reply(b,
			strings.Join(responseField.Response, "\n"),
			&gotgbot.SendMessageOpts{
				ParseMode: "MarkdownV2",
			},
		)
		if err != nil {
			return fmt.Errorf("get price failed: %s", err)
		}
		return handlers.NextConversationState("advertise")
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

func RegisterAdvertise(b *gotgbot.Bot, ctx *ext.Context) error {
	responseField := languages.Response.Conversations.Advertise

	convertFromPersianDigits := func(s string) string {
		persianDigitsMap := map[string]string{
			"۰": "0",
			"۱": "1",
			"۲": "2",
			"۳": "3",
			"۴": "4",
			"۵": "5",
			"۶": "6",
			"۷": "7",
			"۸": "8",
			"۹": "9",
		}
		convertedStr := strings.Map(func(r rune) rune {
			if replacement, ok := persianDigitsMap[string(r)]; ok {
				return []rune(replacement)[0]
			}
			return r
		}, s)
		return convertedStr
	}

	convertedAscii := convertFromPersianDigits(ctx.EffectiveMessage.Text)
	price, err := strconv.ParseInt(convertedAscii, 10, 64)
	if err != nil {
		_, err := ctx.EffectiveMessage.Reply(b,
			responseField.Failed,
			&gotgbot.SendMessageOpts{
				ParseMode: "MARKDOWNV2",
			},
		)
		if err != nil {
			return fmt.Errorf("register advertise failed: %s", err)
		}
		return handlers.NextConversationState("advertise")
	} else if bot.Configs.PriceLimit[0] <= price && price <= bot.Configs.PriceLimit[1] {
		if user, ok := users[ctx.EffectiveMessage.Chat.Id]; ok {
			user.AdvertisePrice = price
		}

		msg, err := ctx.EffectiveMessage.Reply(b,
			strings.Join(responseField.Response, "\n"),
			&gotgbot.SendMessageOpts{
				ParseMode: "MARKDOWNV2",
			},
		)
		if err != nil {
			return fmt.Errorf("register advertise failed: %s", err)
		}

		time.Sleep(2 * time.Second)
		sendDescription(b, ctx)

		_, err = b.DeleteMessage(msg.Chat.Id, msg.MessageId, nil)
		if err != nil {
			return fmt.Errorf("delete message in register advertise failed: %s", err)
		}

		return handlers.EndConversation()
	}
	printer := message.NewPrinter(language.Persian)
	response := printer.Sprintf(
		responseField.FailedLimit,
		bot.Configs.PriceLimit[0],
		bot.Configs.PriceLimit[1],
	)
	_, err = ctx.EffectiveMessage.Reply(b,
		response,
		&gotgbot.SendMessageOpts{
			ParseMode: "MARKDOWNV2",
		},
	)
	if err != nil {
		return fmt.Errorf("register advertise failed: %s", err)
	}
	return handlers.NextConversationState("advertise")
}

func sendDescription(b *gotgbot.Bot, ctx *ext.Context) error {
	responseField := languages.Response.Conversations.Description
	printer := message.NewPrinter(language.Persian)
	response := printer.Sprintf(strings.Join(responseField.Response, "\n"),
		ctx.EffectiveMessage.Chat.FirstName,
		users[ctx.EffectiveMessage.Chat.Id].AdvertiseDescription,
		users[ctx.EffectiveMessage.Chat.Id].Username,
		users[ctx.EffectiveMessage.Chat.Id].AdvertisePrice,
	)
	_, err := ctx.EffectiveMessage.Reply(b,
		response,
		&gotgbot.SendMessageOpts{
			ParseMode: "MARKDOWNV2",
		},
	)
	if err != nil {
		return fmt.Errorf("send description failed: %s", err)
	}
	return nil
}
