package main

import (
	"currency-quotes-api/internal/service"
	"currency-quotes-api/internal/transport"
)

func main() {
	service := service.NewService()
	handler := transport.NewHandler(service)

	server := transport.NewServer(handler)
	transport.InitRoutes(server)

}
