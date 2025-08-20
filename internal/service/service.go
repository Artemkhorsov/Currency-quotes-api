package service

type Service struct {
	Currency string `json:"currency"`
	Value float64 `json:"value"`
}

type RateService interface {
	AddOrUpdateRate(service Service) error
	GetList(currency string)(Rates error)
	ListRate()([]Service, error)
	DeleteRate(service Service) error
	
}
