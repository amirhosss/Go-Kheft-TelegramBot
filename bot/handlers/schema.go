package handlers

type user struct {
	chatId               int64
	username             string
	advertisePrice       int64
	advertiseDescription string
}

var users map[int64]*user = make(map[int64]*user)
