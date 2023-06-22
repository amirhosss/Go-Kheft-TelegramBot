package handlers

type user struct {
	chatId   int64
	username string
	books    []book
}

type book struct {
	writer      string
	description string
	price       int64
}

var users map[int64]*user = make(map[int64]*user)
