package server

import (
	"lesson15/internal/handlers"
	"lesson15/internal/services"
	"net/http"
)

type Server struct {
	Mux     *http.ServeMux
	Service *services.Service
}

func NewServer(mux *http.ServeMux, service *services.Service) *Server {
	return &Server{
		Mux:     mux,
		Service: service,
	}
}

func (s *Server) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	s.Mux.ServeHTTP(response, request)
}

func (s *Server) Init() {

	s.Mux.HandleFunc("/Calculate", handlers.Calculation)
	s.Mux.HandleFunc("/Gethistory", handlers.GetHistory)
	s.Mux.HandleFunc("/Cleanhistory", handlers.CleanHistory)
}
