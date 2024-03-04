package ingridients

import "context"

type Functionality interface {
	All(ctx context.Context) ([]Ingridient, error)
	Random(ctx context.Context, limit int) ([]Ingridient, error)
	Find(ctx context.Context, search string, limit, last int) ([]Ingridient, error)
}

type Default struct {
	StorageOperator IngridientsOperator
}

func (d *Default) All(ctx context.Context) ([]Ingridient, error) {
	return d.StorageOperator.All(ctx)
}

func (d *Default) Random(ctx context.Context, limit int) ([]Ingridient, error) {
	return d.StorageOperator.Random(ctx, limit)
}

func (d *Default) Find(ctx context.Context, search string, limit, last int) ([]Ingridient, error) {
	return d.StorageOperator.Find(ctx, search, limit, last)
}
