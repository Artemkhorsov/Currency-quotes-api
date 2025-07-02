package transport

type Server struct {
	Handler *Handler
}
type Error interface {
	Error() string
}

func NewServer(handler Handler) *Server {
	return &Server{
		Handler: &handler,
	}
}
