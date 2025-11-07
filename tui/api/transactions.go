package api

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TransactionType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetTransactionTypes() ([]TransactionType, error) {
	transactions := []TransactionType{}

	res, err := Request("GET", "/transactions/", nil)

	if err != nil {
		logger.Logf("%s", err.Error())
		return transactions, err
	}

	if !res.IsStatusSuccess() {
		err = errors.New(
			fmt.Sprintf(
				"Unable to retrieve transaction types. %s. %s",
				res.Res().Status,
				string(res.Body()),
			),
		)

		logger.Logf("%s", err.Error())
		return transactions, err
	}

	res.Unmarshall(&transactions)

	return transactions, nil
}

type Transaction struct {
	Amount  int    `json:"amount"`
	Type    int    `json:"type"`
	Comment string `json:"comment"`
}

func CreateTransaction(amount int, transactionType int, comment string) error {
	jsonString, _ := json.Marshal(
		Transaction{
			Amount:  amount,
			Type:    transactionType,
			Comment: comment,
		},
	)

	res, err := Request("PUT", "/transactions/", jsonString)

	if err != nil {
		logger.Logf("%s", err.Error())
		return err
	}

	if !res.IsStatusSuccess() {
		err = errors.New(
			fmt.Sprintf(
				"Unable to create an transaction. %s. %s",
				res.Res().Status,
				string(res.Body()),
			),
		)

		logger.Logf("%s", err.Error())
		return err
	}

	return nil
}
