package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Server struct {
	Mux  *http.ServeMux
	Conn *pgx.Conn
	Port int
}

func NewServer(mux *http.ServeMux, conn *pgx.Conn, port int) Server {
	return Server{
		Mux:  mux,
		Conn: conn,
		Port: port,
	}
}

func (s *Server) Start() {
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Mux)

	if err != nil {
		log.Fatal(err.Error())
	}
}
