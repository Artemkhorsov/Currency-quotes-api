package transport

import (
	"fmt"
	"log"
	"net/http"
)

func InitRoutes(s *Server) {
	http.HandleFunc("/rate", s.Handler.AddOrUpdateRateHandler)

	http.HandleFunc("/delete", s.Handler.DeleteHandler)

	http.HandleFunc("/rates", s.Handler.GetList)

	http.HandleFunc("/convert", s.Handler.ConvertTheAmount)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
