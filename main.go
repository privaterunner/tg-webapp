package main

import (
	"fmt"

	tgbotapi "gitlab.com/kingofsystem/telegram-bot-api/v5"
)

func main() {
	bot, _ := tgbotapi.NewBotAPI("5876584894:AAEQ6LDBmt07uBn4sDVoJn31GPDUI_D2qac")
	up := tgbotapi.NewUpdate(0)
	up.Timeout = 60

	updates := bot.GetUpdatesChan(up)
	for update := range updates {
		fmt.Println("bot started")
		switch {
		case update.Message != nil:
			if update.Message.Text == "/start" {
				fmt.Println("hh", update.Message.From.ID)
				msg := tgbotapi.NewMessage(update.Message.From.ID, "hello")
				msg.ReplyMarkup = markup()
				bot.Send(msg)
			}
		}
	}
}

func markup() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonWebApp("Let's Play 🎮", tgbotapi.WebAppInfo{
				URL: "https://tg-webapp.onrender.com/",
			}),
		),
	)
}
