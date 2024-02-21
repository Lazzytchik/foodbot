package ingridients

type Ingridient struct {
	ID   int    `db:"id"`
	Name string `db:"name"`

	Calories   float32 `db:"calories"`
	Nutritions float32 `db:"nutritions"`
	Fats       float32 `db:"fats"`
	Carbs      float32 `db:"carbs"`
	Water      float32 `db:"water"`
	Fibers     float32 `db:"fiber"`
}
