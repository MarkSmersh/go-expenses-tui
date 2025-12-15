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

type TransactionExtended struct {
	Transaction
	ID        int    `json:"id"`
	CreatedAt int    `json:"created_at"`
	TypeName  string `json:"type_name"`
}

type GetTransactionsReq struct {
	Count int `json:"count"`
	From  int `json:"from"`
	To    int `json:"to"`
	Type  int `json:"type"`
}

func GetTransactions(count int, from int, to int, transactionType int) ([]TransactionExtended, error) {
	data := GetTransactionsReq{
		Count: count,
		From:  from,
		To:    to,
		Type:  transactionType,
	}

	json, _ := json.Marshal(data)

	body := []TransactionExtended{}

	res, err := Request(
		"POST",
		"/transactions/",
		json,
	)

	if err != nil {
		logger.Logf(
			"Unable to get transactions. %s",
			err.Error(),
		)
		return body, err
	}

	res.Unmarshall(&body)

	return body, nil
}

func DeleteTransaction(id int) error {
	res, err := Request(
		"DELETE",
		fmt.Sprintf("/transactions/%d", id),
		nil,
	)

	if err != nil {
		logger.Logf("Unable to delete the transaction. %s", err.Error())
		return err
	}

	if !res.IsStatusSuccess() {
		err = errors.New(
			fmt.Sprintf("%s. %s", res.Res().Status, string(res.Body())),
		)

		logger.Logf("Unable to delete the transaction. %s", err.Error())
		return err
	}

	return nil
}
