package bot

type configs struct {
	BotToken          string
	ChannelChatId     int64
	ChannelUsername   string
	RegistrationPrice int64
}

var Configs *configs = &configs{
	BotToken:          "6092617943:AAEz3506OD7ni_O796972eyOn6-OnJiC4Pc",
	ChannelChatId:     -1001202920496,
	ChannelUsername:   "@kheft\\_channel",
	RegistrationPrice: 5000,
}
