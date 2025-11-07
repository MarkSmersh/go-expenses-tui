package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarkSmersh/go-expenses-tui.git/api/components"
	"github.com/MarkSmersh/go-expenses-tui.git/api/models"
)

func (s *Server) TransactionsRouter(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access-token")

	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	}

	user := models.NewUser(s.Conn)
	apierr := user.Auth(cookie.Value)

	if apierr != nil {
		w.WriteHeader(apierr.Code())
		w.Write(apierr.ErrorBytes())
		return
	}

	username, apierr := user.GetUsername()

	if apierr != nil {
		w.WriteHeader(apierr.Code())
		w.Write(apierr.ErrorBytes())
		return
	}

	switch r.Method {
	case "GET":
		s.TransactionsGet(w, r, username)
	case "PUT":
		s.TransactionsPut(w, r, username)
	case "DELETE":
		s.TransactionsDelete(w, r, username)
	}
}

// Get transaction types within their ids and names
func (s *Server) TransactionsGet(w http.ResponseWriter, r *http.Request, username string) {
	t := models.NewTransaction(s.Conn)
	transactionTypes, err := t.GetTransactionTypes()

	if err != nil {
		w.WriteHeader(err.Code())
		w.Write(err.ErrorBytes())
		return
	}

	jsonTransactionTypes, jsonerr := json.Marshal(transactionTypes)

	if jsonerr != nil {
		w.WriteHeader(500)
		w.Write(components.InternalServerError().ErrorBytes())
	}

	w.WriteHeader(200)
	w.Write(jsonTransactionTypes)
}

func (s *Server) TransactionsPut(w http.ResponseWriter, r *http.Request, username string) {
	data, err := models.NewTransactionData(r.Body)

	if err != nil {
		w.WriteHeader(err.Code())
		w.Write(err.ErrorBytes())
		return
	}

	transaction := models.NewTransaction(s.Conn)

	err = transaction.Create(data.Amount, data.Comment, data.Type, username)

	if err != nil {
		w.WriteHeader(err.Code())
		w.Write(err.ErrorBytes())
		return
	}

	w.WriteHeader(201)
	w.Write([]byte("Transaction has been created succesfully"))
}

func (s *Server) TransactionsDelete(w http.ResponseWriter, r *http.Request, username string) {
	splittedUrl := strings.Split(r.URL.Path, "/")
	slugString := splittedUrl[len(splittedUrl)-1]

	slug, err := strconv.Atoi(slugString)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Slug must be int value"))
		return
	}

	transaction := models.NewTransaction(s.Conn)

	apierr := transaction.Delete(slug, username)

	if apierr != nil {
		w.WriteHeader(apierr.Code())
		w.Write(apierr.ErrorBytes())
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Transaction has been succesfully deleted"))
}
