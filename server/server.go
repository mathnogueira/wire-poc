package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	ht "github.com/mathnogueira/ioc/http"
)

type Server struct {
	router     *mux.Router
	controller ht.Controller
}

func NewServer(controller ht.Controller) *Server {
	s := &Server{
		router:     mux.NewRouter(),
		controller: controller,
	}

	s.setRoutes()

	return s
}

func (s *Server) setRoutes() {
	s.router.HandleFunc("/", s.controller.Greet).Methods("GET")
}

func (s *Server) Start() {
	srv := &http.Server{
		Handler: s.router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
