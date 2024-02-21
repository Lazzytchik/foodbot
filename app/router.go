package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler func(tgbotapi.Update)

type BotRouter struct {
	routes map[string]BotHandler
	parser QueryParser
}

type Params struct {
	Path   string
	Params []string
}

func NewBotRouter(srv *Service) *BotRouter {
	routes := map[string]BotHandler{
		"/ingridients":      srv.AllIngridients,
		"/randomIngridient": srv.RandomIngridient,
	}

	router := &BotRouter{
		routes: routes,
		parser: &DefaultParser{},
	}

	return router
}

func (r *BotRouter) Handle(query string, update tgbotapi.Update) {
	params := r.parser.Parse(query)
	handler, valid := r.routes[params.Path]
	if !valid {
		return
	}

	handler(update)
}
