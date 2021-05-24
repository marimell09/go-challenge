package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/marimell09/stone-challenge/db"
	"github.com/marimell09/stone-challenge/models"
	"github.com/marimell09/stone-challenge/utils"
)

var account_id_key = "accountId"

func accounts(router chi.Router) {
	router.Get("/", getAllAccounts)
	router.Post("/", createAccount)
	router.Route("/{accountId}", func(router chi.Router) {
		router.Use(AccountContext)
		router.Get("/", getAccount)
		router.Put("/", updateAccount)
		router.Delete("/", deleteAccount)
		router.Get("/balance", getAccountBalance)
	})
}

//Account context responsible to save the account id
func AccountContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		account_id := chi.URLParam(r, "accountId")
		if account_id == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Account Id is required")))
			return
		}

		id, err := strconv.Atoi(account_id)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("Invalid Account Id")))
		}

		ctx := context.WithValue(r.Context(), account_id_key, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//Create account, registering the secret in hash format
func createAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	if err := render.Bind(r, account); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	account.Secret = utils.HashSecret(account.Secret)

	if err := dbInstance.AddAccount(account); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	account.Secret = `*********`
	if err := render.Render(w, r, account); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

//Get all accounts
func getAllAccounts(w http.ResponseWriter, r *http.Request) {
	accounts, err := dbInstance.GetAllAccounts()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, accounts); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

//Get account for the specified account id in context
func getAccount(w http.ResponseWriter, r *http.Request) {
	account_id := r.Context().Value(account_id_key).(int)
	account, err := dbInstance.GetAccountById(account_id)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &account); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

//Delete account based on the given account id in context
func deleteAccount(w http.ResponseWriter, r *http.Request) {
	account_id := r.Context().Value(account_id_key).(int)
	err := dbInstance.DeleteAccount(account_id)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

//Update account information, generating new secret in hash
func updateAccount(w http.ResponseWriter, r *http.Request) {
	account_id := r.Context().Value(account_id_key).(int)
	account_data := models.Account{}

	if err := render.Bind(r, &account_data); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	account_data.Secret = utils.HashSecret(account_data.Secret)

	account, err := dbInstance.UpdateAccount(account_id, account_data)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	account.Secret = `*********`
	if err := render.Render(w, r, &account); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

//Get account balance for the given account id in context
func getAccountBalance(w http.ResponseWriter, r *http.Request) {
	account_id := r.Context().Value(account_id_key).(int)
	balance, err := dbInstance.GetAccountBalanceById(account_id)

	res := struct {
		Balance float64 `json:"balance"`
	}{balance}

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}

	render.JSON(w, r, res)
}
