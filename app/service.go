package app

import (
	"context"
	"log"

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

func (s *Service) AllIngridients(update tgbotapi.Update, params Params) error {
	ing, err := s.Ingridients.All(context.Background())
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Невозможно получить ингридиенты")))
		return err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		ingridients.StringPresenter{Model: ing}.String(),
	)
	_, err = s.Bot.Send(msg)
	return err
}

func (s *Service) RandomIngridient(update tgbotapi.Update, params Params) error {
	var limit = 1

	err := params.TouchParam(0)
	if err == nil {
		limit, err = params.GetInt(0)
		if err != nil {
			s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Некорректный параметр")))
			return err
		}
	}

	log.Println("dude")

	ing, err := s.Ingridients.Random(context.Background(), limit)
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Невозможно получить ингридиенты")))
		return err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		ingridients.StringPresenter{Model: ing}.String(),
	)
	_, err = s.Bot.Send(msg)
	return err
}

func (s *Service) FindIngridient(update tgbotapi.Update, params Params) error {
	var searchString string

	var limit = 5
	var last = 0

	err := params.TouchParam(0)
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Некорректный параметр")))
		return err
	}

	searchString, err = params.GetString(0)
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Некорректный параметр")))
		return err
	}

	ings, err := s.Ingridients.Find(context.Background(), searchString, limit, last)
	if err != nil {
		s.SendErrorMessage(update, s.Errors.Error(err, errors.New("Невозможно найти ингридиент")))
		return err
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вот ингридиенты по запросу "+searchString+":")
	msg.ReplyMarkup = ingridients.InlineKeyboardPresenter{Model: ings}.Keyboard()

	_, err = s.Bot.Send(msg)
	return err
}

func (s *Service) SendErrorMessage(update tgbotapi.Update, err error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	s.Bot.Send(msg)
}
