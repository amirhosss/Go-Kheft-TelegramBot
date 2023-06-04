package configs

type configs struct {
	BotToken        string
	ChannelChatId   int64
	ChannelUsername string
}

var Configs *configs = &configs{
	BotToken:        "6092617943:AAGPJlTtZpNwKjxmy4fF5x0i9-62OThyUpg",
	ChannelChatId:   -1001202920496,
	ChannelUsername: "@kheft\\_channel",
}
