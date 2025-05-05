package api

import (
	"TunaAPIGateway/internal/api/routers"
	"log"
	"net/http"
)

type Server struct {
	addr string
}

func NewAPIServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	router.Handle("/events/", http.StripPrefix("/events", routers.EventsRouter()))
	log.Println("Starting server on port", s.addr)

	gatewayHandler := AccessLogMiddleware(router)
	gatewayHandler = PanicMiddleware(gatewayHandler)

	return http.ListenAndServe(s.addr, gatewayHandler)
}
