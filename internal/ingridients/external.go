package ingridients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type External struct {
	URL string

	Client http.Client
}

func (d *External) All(ctx context.Context) ([]Ingridient, error) {
	url := fmt.Sprintf("http://%s/api/v1/ingridients", d.URL)

	resp, err := d.Client.Get(url)
	if err != nil {
		return []Ingridient{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []Ingridient{}, fmt.Errorf("Error status code %d: %s", resp.StatusCode, resp.Body)
	}

	var ingridients []Ingridient
	err = json.NewDecoder(resp.Body).Decode(&ingridients)

	return ingridients, err
}

func (d *External) Random(ctx context.Context, limit int) ([]Ingridient, error) {
	url := fmt.Sprintf("http://%s/api/v1/ingridients/random?limit=%d", d.URL, limit)

	resp, err := d.Client.Get(url)
	if err != nil {
		return []Ingridient{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []Ingridient{}, fmt.Errorf("Error status code %d: %s", resp.StatusCode, resp.Body)
	}

	var ingridients []Ingridient
	err = json.NewDecoder(resp.Body).Decode(&ingridients)

	return ingridients, err
}

// is it find though
func (d *External) Find(ctx context.Context, search string, limit, last int) ([]Ingridient, error) {
	url := fmt.Sprintf("http://%s/api/v1/ingridients/find?limit=%d&last=%d&search=%s", d.URL, limit, last, search)

	log.Println(url)

	resp, err := d.Client.Get(url)
	if err != nil {
		return []Ingridient{}, err
	}

	if resp.StatusCode != http.StatusOK {
		text, _ := io.ReadAll(resp.Body)
		return []Ingridient{}, fmt.Errorf("Error status code %d: %s", resp.StatusCode, string(text))
	}

	var ingridients []Ingridient
	err = json.NewDecoder(resp.Body).Decode(&ingridients)

	return ingridients, err
}
