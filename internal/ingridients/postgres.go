package ingridients

import "context"

func (p *Postgres) All(ctx context.Context) (data []Ingridient, err error) {
	q := `
		select * from ingridients
	`

	err = p.DB.SelectContext(ctx, &data, q)
	return
}
