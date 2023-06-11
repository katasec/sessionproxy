package core

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	config *Config
	router *http.ServeMux
}

func NewServer(config ...*Config) *Server {

	var c *Config
	if config == nil {
		c = NewConfig()
	} else {
		c = config[0]
	}

	return &Server{
		config: c,
		router: http.NewServeMux(),
	}

}

func (s *Server) Start() {
	s.initaliseRoutes()

	addr := fmt.Sprintf("0.0.0.0:%s", s.config.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	log.Printf("Starting server on: http://%s\n", addr)
	log.Fatal(server.ListenAndServe())
}
