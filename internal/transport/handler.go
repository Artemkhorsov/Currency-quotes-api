package transport

import (
	convert "currency-quotes-api/internal/core"
	"currency-quotes-api/internal/core/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	service Service
}

type Service interface {
	AddOrUpdateRate(rates models.Rates) error
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) AddOrUpdateRateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var ratesStruct models.Rates
	err := json.NewDecoder(r.Body).Decode(&ratesStruct)
	if err != nil {
		log.Println("Ошибка декодирования JSON", err)
		http.Error(w, "Неверный формат", http.StatusBadRequest)

		return
	}
	err = h.service.AddOrUpdateRate(ratesStruct)

	log.Printf("Валюта '%s' успешно добавлена", ratesStruct.Currency)

	if err != nil {
		log.Println("Ошибочка при добавлении/обновлении курса", err)
		http.Error(w, "Не удалось добавить/обновить курс", http.StatusInternalServerError)
	}

}

func (h Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var deleteRequest models.DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&deleteRequest)
	if err != nil {
		log.Println("Ошибка декодирования JSON", err)
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}
	currencyKey := deleteRequest.Currency

	if _, ok := models.CurrencyNamesMap[currencyKey]; !ok {
		msg := fmt.Sprintf("Валюта '%s' не найдена", currencyKey)
		log.Println(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	delete(models.CurrencyNamesMap, currencyKey)

	log.Printf("Валюта '%s' успешно удалена", currencyKey)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Удаление выполнено успешно!"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println("Ошибка маршалинга JSON", err)
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResponse)
}

func (h Handler) GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	err := json.NewEncoder(w).Encode(models.CurrencyNamesMap)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h Handler) ConvertTheAmount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Ошибка декодирования JSON ", err)
		http.Error(w, "Ошибка конвертации ", http.StatusBadRequest)
		return
	}
	amount := req.Amount
	fromCurrency := req.FromCurrency
	toCurrency := req.ToCurrency

	convertTheAmount, err := convert.ConvertsTheRate(amount, fromCurrency, toCurrency)
	if err != nil {
		log.Printf("ошибка конвертации:%v", err)
		http.Error(w, "ошибка конвертации", http.StatusInternalServerError)
		return
	}
	resp := models.Response{
		ConvertTheAmount: convertTheAmount,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Println("Ошибка маршалинга JSON:", err)
		http.Error(w, "Ошибка маршалинга JSON", http.StatusInternalServerError)
		return

	}
}
