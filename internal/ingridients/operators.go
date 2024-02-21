package ingridients

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IngridientsOperator interface {
	All(ctx context.Context) ([]Ingridient, error)
}

type Postgres struct {
	DB *sqlx.DB
}