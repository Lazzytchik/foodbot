package ingridients

import (
	"fmt"
	"strings"
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
