package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewConn(uri string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), uri)

	if err != nil {
		log.Fatal(err.Error())
	}

	return conn
}
