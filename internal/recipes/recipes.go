package recipes

import (
	"time"

	"github.com/lazzytchik/foodbot/internal/ingridients"
)

type Recipe struct {
	Name string

	Calories float64

	Nutritions float64
	Carbs      float64
	Fats       float64

	Water float64
	Fiber float64

	CookTime time.Duration

	Steps       []CookStep
	Ingridients []ingridients.Ingridient
}

type CookStep struct {
	Title   string
	Content string
	Order   int

	RecipeId int
}
