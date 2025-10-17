package tables

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func TransactionsTable(conn *pgx.Conn) {
	query := `create table if not exists transactions (
		id BIGSERIAL PRIMARY KEY,
		amount int,
		comment varchar(256),
		transaction_type_id bigserial references transaction_types(id),
		user_id bigserial references users(id),
		created_at timestamp default now()
	)`

	_, err := conn.Exec(context.Background(), query)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
