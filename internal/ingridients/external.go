package ingridients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type External struct {
	URL string

	Client http.Client
}

func (d *External) All(ctx context.Context) ([]Ingridient, error) {
	url := fmt.Sprintf(d.URL + "/ingridients")

	resp, err := d.Client.Get(url)
	if err != nil {
		return []Ingridient{}, err
	}

	var ingridients []Ingridient
	err = json.NewDecoder(resp.Body).Decode(&ingridients)

	return ingridients, err
}

func (d *External) Random(ctx context.Context, limit int) ([]Ingridient, error) {
	url := fmt.Sprintf(d.URL+"/ingridients/random?limit=%d", limit)

	resp, err := d.Client.Get(url)
	if err != nil {
		return []Ingridient{}, err
	}

	var ingridients []Ingridient
	err = json.NewDecoder(resp.Body).Decode(&ingridients)

	return ingridients, err
}

func (d *External) Find(ctx context.Context, search string, limit, last int) ([]Ingridient, error) {
	url := fmt.Sprintf(d.URL+"/ingridients/random?limit=%d&last=%d&search=%s", limit, last, search)

	resp, err := d.Client.Get(url)
	if err != nil {
		return []Ingridient{}, err
	}

	var ingridients []Ingridient
	err = json.NewDecoder(resp.Body).Decode(&ingridients)

	return ingridients, err
}
