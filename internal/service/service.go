package service

import "currency-quotes-api/internal/core/models"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s Service) AddOrUpdateRate(rates models.Rates) error {
	models.CurrencyNamesMap[rates.Currency] = struct{}{}

	return nil
}
