package ingridients

type Ingridient struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`

	Calories   float32 `db:"calories" json:"calories"`
	Nutritions float32 `db:"nutritions" json:"nutritions"`
	Fats       float32 `db:"fats" json:"fats"`
	Carbs      float32 `db:"carbs" json:"carbs"`
	Water      float32 `db:"water" json:"water"`
	Fibers     float32 `db:"fibers" json:"fibers"`
}
