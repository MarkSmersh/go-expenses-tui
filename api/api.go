package api

import (
	"log"
	"net/http"
	"os"

	"github.com/MarkSmersh/go-expenses-tui.git/db"
	"github.com/MarkSmersh/go-expenses-tui.git/tables"
)

func Init() {
	mux := http.NewServeMux()

	dburi := os.Getenv("DB_URI")

	if dburi == "" {
		log.Fatal("Enviroment variable DB_URI is absent!")
	}

	conn := db.NewConn(dburi)
	tables.Init(conn)

	server := NewServer(mux, conn, 1488)
	server.Mux.HandleFunc("/transactions/", server.TransactionsRouter)
	server.Mux.HandleFunc("/auth/", server.AuthRouter)

	log.Print("Server has started successfully!")

	server.Start()
}
