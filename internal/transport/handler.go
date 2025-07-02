package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) AddOrUpdateRateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var rates Rates
	err := json.NewDecoder(r.Body).Decode(&rates)
	if err != nil {
		log.Println("Ошибочка", err)
		http.Error(w, "Неверный формат", http.StatusBadRequest)

		return
	}
	fmt.Println(rates.Currency, rates.Value)
}

func (h Handler) GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	js, err := json.Marshal(GetCurrency)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(GetCurrency)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func (h Handler) ConvertTheAmount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)

		return
	}
}
