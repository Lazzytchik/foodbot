package ingridients

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IngridientsOperator interface {
	All(ctx context.Context) ([]Ingridient, error)
	Random(ctx context.Context, limit int) ([]Ingridient, error)
	Find(ctx context.Context, search string, limit, last int) (data []Ingridient, err error)
}

type Postgres struct {
	DB *sqlx.DB
}
