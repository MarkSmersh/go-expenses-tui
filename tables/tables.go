package tables

import (
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func Init(conn *pgx.Conn) {
	TransactionTypesTable(conn)
	UsersTable(conn)
	TransactionsTable(conn)

	slog.Debug("SQL Tables defined")
}
