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
)

var accountIdKey = "accountId"

func accounts(router chi.Router) {
	router.Get("/", getAllAccounts)
	router.Post("/", createAccount)
	router.Route("/{accountId}", func(router chi.Router) {
		router.Use(AccountContext)
		router.Get("/", getAccount)
		router.Put("/", updateAccount)
		router.Delete("/", deleteAccount)
	})
}

func AccountContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accountId := chi.URLParam(r, "accountId")
		if accountId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("account ID is required")))
			return
		}

		id, err := strconv.Atoi(accountId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid account ID")))
		}

		ctx := context.WithValue(r.Context(), accountIdKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	if err := render.Bind(r, account); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddAccount(account); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, account); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

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

func getAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value(accountIdKey).(int)
	account, err := dbInstance.GetAccountById(accountId)
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

func deleteAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value(accountIdKey).(int)
	err := dbInstance.DeleteAccount(accountId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.Context().Value(accountIdKey).(int)
	accountData := models.Account{}

	if err := render.Bind(r, &accountData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}

	account, err := dbInstance.UpdateAccount(accountId, accountData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &account); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
