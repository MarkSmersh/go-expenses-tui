package tables

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

// Should I create default transaction_types like for food, transport etc?
func TransactionTypesTable(conn *pgx.Conn) {
	query := `create table if not exists transaction_types (
		id BIGSERIAL PRIMARY KEY,
		name varchar(64) UNIQUE
	);`

	_, err := conn.Exec(context.Background(), query)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
