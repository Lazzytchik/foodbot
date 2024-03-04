package ingridients

import "context"

func (p *Postgres) All(ctx context.Context) (data []Ingridient, err error) {
	q := `
		select * from ingridients
	`

	err = p.DB.SelectContext(ctx, &data, q)
	return
}

func (p *Postgres) Random(ctx context.Context, limit int) (data []Ingridient, err error) {
	q := `
		select * from ingridients
		order by random()
		limit $1
	`

	err = p.DB.SelectContext(ctx, &data, q, limit)
	return
}

func (p *Postgres) Find(ctx context.Context, search string, limit int, last int) (data []Ingridient, err error) {
	q := `
		select
			 *
		from ingridients i 
		where 
			to_tsvector('russian', name) @@ to_tsquery('russian', $1)
			and id > $3
		limit $2
	`

	err = p.DB.SelectContext(ctx, &data, q, search, limit, last)
	return
}
