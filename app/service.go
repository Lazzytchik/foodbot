package app

import (
	"context"

	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	e "github.com/lazzytchik/foodbot/internal/errors"
	"github.com/lazzytchik/foodbot/internal/ingridients"
)

type Service struct {
	Ingridients ingridients.Functionality

	Errors e.Factory
	Bot    *tgbotapi.BotAPI
	Router *BotRouter
}

func (s *Service) AllIngridients(update tgbotapi.Update) {
	ing, err := s.Ingridients.All(context.Background())
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Невозможно получить ингридиенты")))
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		ingridients.StringPresenter{Model: ing}.String(),
	)
	s.Bot.Send(msg)
}

func (s *Service) RandomIngridient(update tgbotapi.Update) {
	ing, err := s.Ingridients.Random(context.Background(), 1)
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Невозможно получить ингридиенты")))
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		ingridients.StringPresenter{Model: ing}.String(),
	)
	s.Bot.Send(msg)
}

func (s *Service) SendErrorMessage(update tgbotapi.Update, err error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	s.Bot.Send(msg)
}
