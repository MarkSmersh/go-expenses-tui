package models

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"github.com/MarkSmersh/go-expenses-tui.git/api/components"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type TransactionData struct {
	Amount  int    `json:"amount"`
	Comment string `json:"comment"`
	Type    int    `json:"type"`
}

func NewTransactionData(body io.ReadCloser) (TransactionData, components.ApiError) {
	var data TransactionData

	decoder := json.NewDecoder(body)
	// decoder.UseNumber()
	decoder.Decode(&data)

	if data.Amount <= 0 {
		return TransactionData{}, components.NewApiError(400, "Invalid amount for transaction. The value should be more than 0")
	}

	if len(data.Comment) > 256 {
		return TransactionData{}, components.NewApiError(400, "Invalid length for the comment field. The value should have length less than 256 sumbols")
	}

	return data, nil
}

type Transaction struct {
	conn *pgx.Conn
}

func NewTransaction(conn *pgx.Conn) Transaction {
	return Transaction{
		conn: conn,
	}
}

type TransactionType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (t Transaction) GetTransactionTypes() ([]TransactionType, components.ApiError) {

	transactionTypes := []TransactionType{}

	rows, err := t.conn.Query(
		context.Background(),
		"SELECT id, name FROM transaction_types",
	)

	if err != nil {
		slog.Error(err.Error())
		return transactionTypes, components.InternalServerError()
	}

	for i := 0; ; i++ {
		if !rows.Next() {
			break
		}

		var id int
		var name string

		err := rows.Scan(&id, &name)

		if err != nil {
			slog.Error(err.Error())
			return transactionTypes, components.InternalServerError()
		}

		transactionTypes = append(
			transactionTypes,
			TransactionType{
				ID:   id,
				Name: name,
			},
		)
	}

	return transactionTypes, nil
}

func (t Transaction) Create(amount int, comment string, transactionType int, username string) components.ApiError {
	_, err := t.conn.Exec(
		context.Background(),
		"INSERT INTO transactions (amount, comment, transaction_type_id, user_id) VALUES ($1, $2, $3, (SELECT id FROM users WHERE username = $4))",
		amount,
		comment,
		transactionType,
		username,
	)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.SQLState() == "23503" {
				return components.NewApiError(400, "One or many options are using unexisting references")
			}

			slog.Error(pgErr.Error())
			return components.InternalServerError()
		}

		slog.Error(pgErr.Error())
		return components.InternalServerError()
	}

	return nil
}

func (t Transaction) Delete(id int, username string) components.ApiError {
	rows, err := t.conn.Query(
		context.Background(),
		"DELETE FROM transactions WHERE id = $1 and user_id = (SELECT id FROM users WHERE username = $2)",
		id,
		username,
	)

	rows.Close()

	if err != nil {
		slog.Error(err.Error())
		return components.InternalServerError()
	}

	if rows.CommandTag().RowsAffected() <= 0 {
		return components.NewApiError(404, "Nothing to delete")
	}

	return nil
}

type TransactionDataWithTypes struct {
	TransactionData
	ID        int    `json:"id"`
	CreatedAt int    `json:"created_at"`
	TypeName  string `json:"type_name"`
}

func (t Transaction) Find(
	limit int,
	from int,
	to int,
	transactionType int,
	username string,
) ([]TransactionDataWithTypes, components.ApiError) {
	transactions := []TransactionDataWithTypes{}

	rows, err := t.conn.Query(
		context.Background(),
		`select transactions.id, transactions.amount, transactions.comment,
			extract(epoch from transactions.created_at)::integer,
			transaction_types.id as type_id, transaction_types.name as type_name
		from transactions
		join transaction_types on transactions.transaction_type_id = transaction_types.id
		where created_at between to_timestamp($1) and (case when $2 <= 0 then now() else to_timestamp($2) end)
			and user_id = (select id from users where username = $3)
			and (case when $4 > 0 then transaction_types.id = $4 else true end)
		order by transactions.id DESC
		limit $5`,
		from,
		to,
		username,
		transactionType,
		limit,
	)

	if err != nil {
		slog.Error(err.Error())
		return transactions, components.InternalServerError()
	}

	for i := 0; ; i++ {
		if !rows.Next() {
			break
		}

		var id, amount, createdAt, typeId int
		var comment, typeName string

		err := rows.Scan(&id, &amount, &comment, &createdAt, &typeId, &typeName)

		if err != nil {
			slog.Error(err.Error())
			return transactions, components.InternalServerError()
		}

		transactions = append(transactions, TransactionDataWithTypes{
			ID: id,
			TransactionData: TransactionData{
				Amount:  amount,
				Comment: comment,
				Type:    typeId,
			},
			TypeName:  typeName,
			CreatedAt: createdAt,
		})
	}

	return transactions, nil
}
