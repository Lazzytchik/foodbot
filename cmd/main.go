package main

import (
	"log"
	"net/http"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lazzytchik/foodbot/app"
	"github.com/lazzytchik/foodbot/internal/errors"
	"github.com/lazzytchik/foodbot/internal/ingridients"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("foodbot")
	token := "6892396673:AAGliGYDDux8jDMNoZjqpxqnOE6pT5en_Ss"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Bot API init error: ", err)
	}

	service := &app.Service{
		Ingridients: &ingridients.External{
			URL: os.Getenv("FOODAPI_URL"),
			Client: http.Client{
				Timeout: time.Minute,
			},
		},
		Errors: errors.Factory{
			Handler: func(err error) {
				log.Println(err)
			},
		},
		Bot: bot,
	}

	service.Router = app.NewBotRouter(service)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil {
			log.Println(update.CallbackQuery)
			err := service.Router.Handle(update.CallbackQuery.Data, update)
			if err != nil {
				log.Println(err)
			}
		}
		if update.Message != nil {
			err := service.Router.Handle(update.Message.Text, update)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
