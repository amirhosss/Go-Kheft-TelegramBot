# Kheft Telegrambot

This bot is a client for communicating between users on the Telegram platform.

Users place their ad through the bot in a designated channel so that people who intend to buy can buy the book they want.

## Getting Started
- First clone this repo using below command
```bash
git clone https://github.com/amirhosss/Kheft-telegrambot.git
```
- Cd to bot directory
```bash
cd Kheft-telegrambot/bot
```
- Change `configs.default` to `configs.go`
```bash
mv configs.default configs.go
```
- Then initialize `configs.go` fields with your own data
- Finally run project
```bash
cd ..
go run .
```