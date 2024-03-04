package app

import (
	"errors"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler func(tgbotapi.Update, Params) error

type BotRouter struct {
	routes map[string]BotHandler
	parser QueryParser
}

type Params struct {
	Path   string
	Params []string
}

func (p Params) GetInt(index int) (int, error) {
	if err := p.TouchParam(index); err != nil {
		return 0, errors.New("No param")
	}

	return strconv.Atoi(p.Params[index])
}

func (p Params) GetString(index int) (string, error) {
	if err := p.TouchParam(index); err != nil {
		return "", errors.New("No param")
	}

	return p.Params[index], nil
}

func (p Params) TouchParam(index int) error {
	if len(p.Params) <= index {
		return errors.New("No param")
	}

	return nil
}

func NewBotRouter(srv *Service) *BotRouter {
	routes := map[string]BotHandler{
		"/ingridients":      srv.AllIngridients,
		"/randomIngridient": srv.RandomIngridient,
		"/find":             srv.FindIngridient,
	}

	router := &BotRouter{
		routes: routes,
		parser: &DefaultParser{},
	}

	return router
}

func (r *BotRouter) Handle(query string, update tgbotapi.Update) error {
	log.Println(query)

	params := r.parser.Parse(query)
	handler, valid := r.routes[params.Path]
	if !valid {
		return errors.New("Нет такой команды")
	}

	return handler(update, params)
}
