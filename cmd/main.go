package main

import (
	"currency-quotes-api/internal/transport"
)

func main() {
	handler := transport.NewHandler()

	server := transport.NewServer(handler)
	transport.InitRoutes(server)
}
