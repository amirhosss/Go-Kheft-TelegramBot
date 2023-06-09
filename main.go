package main

import (
	"log"
	"net/http"
	"time"

	"kheft/bot"
	myhandlers "kheft/bot/handlers"
	"kheft/bot/languages"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
)

// This bot is as basic as it gets - it simply repeats everything you say.
func main() {
	// Create bot from environment value.
	b, err := gotgbot.NewBot(bot.Configs.BotToken, &gotgbot.BotOpts{
		Client: http.Client{},
		DefaultRequestOpts: &gotgbot.RequestOpts{
			Timeout: gotgbot.DefaultTimeout,
			APIURL:  gotgbot.DefaultAPIURL,
		},
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(&ext.UpdaterOpts{
		Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
			// If an error is returned by a handler, log it and continue going.
			Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
				log.Println("an error occurred while handling update:", err.Error())
				return ext.DispatcherActionNoop
			},
			MaxRoutines: ext.DefaultMaxRoutines,
		}),
	})
	dispatcher := updater.Dispatcher

	registrationHandler := handlers.NewMessage(
		(&bot.CheckMembershipOpts{
			MessageText: languages.Response.Messages.Member.Btns[0],
		}).CheckMessage(b),
		myhandlers.Registration,
	)
	nonMemberStartHandler := handlers.NewMessage(
		(&bot.CheckMembershipOpts{
			ReverseState: true,
		}).CheckMessage(b),
		myhandlers.NonMemberStart,
	)
	memberStartHandler := handlers.NewMessage(
		(&bot.CheckMembershipOpts{}).CheckMessage(b),
		myhandlers.MemberStart)
	nonMemberCheckingHandler := handlers.NewMessage(message.Equal(languages.Response.Messages.NonMember.Btns[0]),
		myhandlers.NonMemberChecking)

	conversation := []ext.Handler{registrationHandler}
	conversationHandler := handlers.NewConversation(
		conversation, map[string][]ext.Handler{
			"registration": {registrationHandler},
		},
		&handlers.ConversationOpts{
			AllowReEntry: true,
		},
	)

	dispatcher.AddHandler(conversationHandler)
	dispatcher.AddHandler(memberStartHandler)
	dispatcher.AddHandler(nonMemberStartHandler)
	dispatcher.AddHandler(nonMemberCheckingHandler)

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: gotgbot.GetUpdatesOpts{
			Timeout: 1,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
