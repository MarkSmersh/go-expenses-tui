package tables

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func UsersTable(conn *pgx.Conn) {
	query := `create table if not exists users (
		id BIGSERIAL PRIMARY KEY,
		username VARCHAR(32) NOT NULL UNIQUE,
		password VARCHAR(256) NOT NULL,
		refresh_tokens text ARRAY,
		created_at timestamp default now()
	)`

	_, err := conn.Exec(context.Background(), query)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
