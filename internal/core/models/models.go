package models

type Currency struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Rate   float64 `json:"rate"`
}

//	var GetCurrency = map[string]Currency{
//		"USD": {"US dollar", "$", 78.53},
//		"EUR": {"Euro", "€", 92.27},
//		"GBP": {"pound sterling", "£", 107.7},
//		"AUD": {"Australian dollar", "$", 51.44},
//		"JPY": {"Japanese yen", "¥", 0.5459},
//		"THB": {"THB", "฿", 2.41},
//		"KRW": {"South Korean won", "₩", 0.057933},
//	}
var CurrencyNamesMap = map[string]struct{}{
	"USD": {}, "EUR": {}, "GBP": {}, "AUD": {},
}

type Rates struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}
type DeleteRequest struct {
	Currency string `json:"currency"`
}

type Request struct {
	Amount       float64 `json:"amount"`
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
}

type Response struct {
	ConvertTheAmount float64 `json:"convertTheAmount"`
}
