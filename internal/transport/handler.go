package transport

import (
	convert "currency-quotes-api/internal/core"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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

	baseURL := "https://v6.exchangerate-api.com/v6/31720714369722230941a8a1/latest/USD"

	parms := url.Values{}
	parms.Add("from", "USD")
	parms.Add("to", "EUR")
	parms.Add("amount", "100")

	fullURL := baseURL + "?" + parms.Encode()
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Ошибка при выполнение запроса", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Ошибка: Неверный статус код:%v\n", resp.StatusCode)
		http.Error(w, fmt.Sprintf("Неверный статус код : %v", resp.StatusCode), http.StatusBadRequest)

		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка при распаковки json", err.Error())
		http.Error(w, "ошибка при чтении ответа", http.StatusInternalServerError)
		return
	}
	fmt.Println("тело ответа", string(body))

	var converts convert.Convert
	err = json.Unmarshal(body, &converts)
	if err != nil {
		log.Println("Ошибка при распаковки json", err)
		http.Error(w, "ошибка при чтении ответа", http.StatusBadRequest)

		return
	}
	jsonConvert, err := json.Marshal(converts)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonConvert)
	if err != nil {
		log.Println("Ошибка при записи ответа", err)
	}
}
