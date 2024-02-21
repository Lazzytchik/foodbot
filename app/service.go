package app

import (
	"context"

	"errors"

	e "github.com/lazzytchik/foodbot/internal/errors"
	"github.com/lazzytchik/foodbot/internal/ingridients"
)

type Service struct {
	Ingridients ingridients.Functionality

	Errors e.Factory
}

func (s *Service) AllIngridients(ctx context.Context) ([]ingridients.Ingridient, error) {
	ing, err := s.Ingridients.All(ctx)
	if err != nil {
		return []ingridients.Ingridient{}, s.Errors.Error(err, errors.New("Невозможно получить ингридиенты"))
	}

	return ing, err
}
