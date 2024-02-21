package app

import (
	"context"
	"strconv"

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

func (s *Service) AllIngridients(update tgbotapi.Update, params Params) {
	ing, err := s.Ingridients.All(context.Background())
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Невозможно получить ингридиенты")))
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		ingridients.StringPresenter{Model: ing}.String(),
	)
	s.Bot.Send(msg)
}

func (s *Service) RandomIngridient(update tgbotapi.Update, params Params) {
	limit := 1
	if len(params.Params) > 0 {
		limit64, err := strconv.ParseInt(params.Params[0], 10, 64)
		if err != nil {
			s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Некорректный параметр")))
			return
		}
		limit = int(limit64)
	}

	ing, err := s.Ingridients.Random(context.Background(), limit)
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Невозможно получить ингридиенты")))
		return
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
