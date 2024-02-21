package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/lazzytchik/foodbot/app"
	"github.com/lazzytchik/foodbot/internal/errors"
	"github.com/lazzytchik/foodbot/internal/ingridients"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("foodpicker")
	token := "6892396673:AAGliGYDDux8jDMNoZjqpxqnOE6pT5en_Ss"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Bot API init error: ", err)
	}

	db, err := sqlx.Open("postgres", "postgresql://lazzy:1111@192.168.1.102:5432/foodpicker?sslmode=disable&application_name=foodbot")
	if err != nil {
		log.Fatalf("Cannot open db: %v", err)
	}

	service := &app.Service{
		Ingridients: &ingridients.Default{
			StorageOperator: &ingridients.Postgres{
				DB: db,
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
		if update.Message != nil {
			service.Router.Handle(update.Message.Text, update)
		}
	}
}
