package ingridients

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StringPresenter struct {
	Model []Ingridient
}

func (ip StringPresenter) String() string {
	var reply strings.Builder
	for _, ing := range ip.Model {
		reply.WriteString(ing.String() + "\n")
	}

	return reply.String()
}

func (i Ingridient) String() string {
	return fmt.Sprintf("Название: %s, Калории: %.1f, Белки: %.1f, Жиры: %.1f, Углеводы: %.1f, Клетчатка: %.1f, Вода: %.1f.", i.Name, i.Calories, i.Nutritions, i.Fats, i.Carbs, i.Fibers, i.Water)
}

type InlineKeyboardPresenter struct {
	Model []Ingridient
}

func (i InlineKeyboardPresenter) Keyboard() tgbotapi.InlineKeyboardMarkup {
	var keyboard tgbotapi.InlineKeyboardMarkup

	var buttons [][]tgbotapi.InlineKeyboardButton

	for _, ing := range i.Model {
		buttons = append(
			buttons,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(ing.Name, "/ingridient "+strconv.Itoa(ing.ID)),
			),
		)
	}

	keyboard = tgbotapi.NewInlineKeyboardMarkup(buttons...)

	return keyboard
}
