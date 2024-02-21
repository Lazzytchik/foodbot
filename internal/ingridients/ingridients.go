package ingridients

import "context"

type Functionality interface {
	All(ctx context.Context) ([]Ingridient, error)
}

type Default struct {
	StorageOperator IngridientsOperator
}

func (d *Default) All(ctx context.Context) ([]Ingridient, error) {
	return d.StorageOperator.All(ctx)
}
